package runner

import (
	"log"
	"reflect"
	"runtime"
)

func Run(i interface{}, args ...interface{}) {
	// Get the function name
	fName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	log.Printf("Running \"%s\"\n", fName)

	// Convert the interface{} function to a reflect.Value and call it with dynamic arguments
	fnValue := reflect.ValueOf(i)

	// Ensure it's a function
	if fnValue.Kind() != reflect.Func {
		log.Println("Provided argument is not a function")
		return
	}

	// Prepare the arguments as reflect.Values
	fnArgs := make([]reflect.Value, len(args))
	for k, v := range args {
		fnArgs[k] = reflect.ValueOf(v)
	}

	// Call the function
	result := fnValue.Call(fnArgs)

	// If the function returns an error, handle it
	if len(result) > 0 {
		if err, ok := result[len(result)-1].Interface().(error); ok && err != nil {
			log.Printf("Warning while running: %v\n", err)
		}
	}
	log.Println("Run success")
}
