package email

type Body struct {
	Name       string
	Intros     []string
	Dictionary []Map
	Table      Table
	Actions    []Action
	Outros     []string
	Greeting   string
	Signature  string
	Title      string
}
