package handlers

import (
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"strings"

	"github.com/toolvox/utilgo/pkg/fsutil"
	"github.com/toolvox/utilgo/pkg/httputils/middlewares"
	"github.com/toolvox/utilgo/pkg/logs"
)

// TemplateModelGetter is a function type that takes a template name and returns a model (data) for that template.
type TemplateModelGetter func(template string) any

// FSTemplateHandler is a struct that handles template rendering from a file system.
type FSTemplateHandler struct {
	fs.FS                         // Embedding file system interface to access templates.
	Patterns  []string            // Selecting templates within the FS.
	FuncMap   template.FuncMap    // Custom template functions.
	GetModel  TemplateModelGetter // Function to get the model for a given template.
	Root      *template.Template  // Root template object.
	DebugMode bool                // Flag to indicate debug mode.
	Log       *slog.Logger        // Logger instance.

	templateExtensions map[string]string // an internal cache of extensions to match file for requests.
}

// NewFSTemplateHandler creates a new instance of FSTemplateHandler with the provided file system and file patterns.
func NewFSTemplateHandler(fs fs.FS, patterns ...string) FSTemplateHandler {
	return FSTemplateHandler{
		FS:       fs,
		Patterns: patterns,
		Root:     template.New("root"),

		templateExtensions: make(map[string]string),
	}
}

// WithModelGetter sets the function to retrieve models for templates.
func (s FSTemplateHandler) WithModelGetter(getter TemplateModelGetter) FSTemplateHandler {
	s.GetModel = getter
	return s
}

// WithModelMap sets a map of models for templates.
func (s FSTemplateHandler) WithModelMap(modelMap map[string]any) FSTemplateHandler {
	s.GetModel = func(template string) any {
		return modelMap[template]
	}
	return s
}

// WithFuncs sets custom functions to be used in templates.
func (s FSTemplateHandler) WithFuncs(funcs template.FuncMap) FSTemplateHandler {
	s.FuncMap = funcs
	s.Root.Funcs(s.FuncMap)
	return s
}

// WithLog sets the logger for the handler.
func (s FSTemplateHandler) WithLog(log *slog.Logger) FSTemplateHandler {
	s.Log = log
	return s
}

// WithDebug enables debug mode.
func (s FSTemplateHandler) WithDebug() FSTemplateHandler {
	s.DebugMode = true
	return s
}

// Parse parses the templates from the file system.
func (s FSTemplateHandler) Parse() FSTemplateHandler {
	s.Root = template.Must(s.Root.ParseFS(s.FS, s.Patterns...))
	for _, pat := range s.Patterns {
		patFiles, err := fsutil.ListFS(s.FS, ".", pat)
		if err != nil {
			s.Log.Error("error listing files", slog.String("pattern", pat), logs.Error(err))
			continue
		}
		for _, file := range patFiles {
			_, name, ext := fsutil.Split(file)
			s.templateExtensions[name] = ext
		}
	}
	return s
}

// ServeHTTP handles HTTP requests and renders templates.
func (s FSTemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	contextLog := r.Context().Value(middlewares.Key_("log"))
	if contextLog != nil {
		s.Log = contextLog.(*slog.Logger)
	}
	if s.DebugMode {
		// Re-parse templates in debug mode for each request.
		s.Root = template.Must(template.New("root").Funcs(s.FuncMap).ParseFS(s.FS, s.Patterns...))
	}

	requestPath := r.URL.Path
	var requestPaths []string
	if strings.Contains(requestPath, "|") {
		// Support for multiple template paths.
		requestPaths = strings.Split(requestPath, "|")
	} else {
		requestPaths = []string{requestPath}
	}

	for _, requestPath := range requestPaths {
		model := s.GetModel(requestPath)
		requestPath = strings.TrimLeft(requestPath, "/")
		tmplName := requestPath + s.templateExtensions[requestPath]
		if model == nil {
			s.Log.Error("missing model for template", slog.String("template", tmplName))
			continue
		}

		tmpl := s.Root.Lookup(tmplName)
		if tmpl == nil {
			s.Log.Error("Failed to lookup template", slog.String("template", tmplName))
			continue
		}

		// Execute the template with the model.
		err := tmpl.Execute(w, model)
		if err != nil {
			s.Log.Error("Failed to execute template", slog.String("template", tmplName), logs.Error(err))
			continue
		}
	}
}
