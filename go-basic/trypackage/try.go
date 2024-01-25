package trypackage

//use capital letter for function name to make it public
//use small letter for function name to make it private

func SayHello(name string) string {
	return "Hello " + name
}

func privateFunction() string {
	return "This is private function"
}

var PublicVariable = "This is public variable"
var privateVariable = "This is private variable"
