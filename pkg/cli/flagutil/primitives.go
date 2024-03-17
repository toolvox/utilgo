package flagutil

type StringValue string

func (sv *StringValue) Set(v string) error { *sv = StringValue(v); return nil }

func (sv StringValue) Get() any { return string(sv) }

func (sv StringValue) String() string { return string(sv) }
