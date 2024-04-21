package evaluator

import (
	"fmt"
	"math"
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"ceil": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.FLOAT_OBJ {
				return newError("argument to `floor` must be ARRAY, got %s",
					args[0].Type())
			}

			f := args[0].(*object.Float)
			ceil := math.Ceil(f.Value)

			return &object.Float{Value: ceil}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"floor": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.FLOAT_OBJ {
				return newError("argument to `floor` must be ARRAY, got %s",
					args[0].Type())
			}

			f := args[0].(*object.Float)
			floor := math.Floor(f.Value)

			return &object.Float{Value: floor}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements)

			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"round": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.FLOAT_OBJ {
				return newError("argument to `floor` must be ARRAY, got %s",
					args[0].Type())
			}

			f := args[0].(*object.Float)
			round := math.Round(f.Value)

			return &object.Float{Value: round}
		},
	},
	"tofloat": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			int, ok := args[0].(*object.Integer)
			if !ok {
				return newError("argument to `tofloat` must be a INTEGER, got %s",
					args[0].Type())
			}

			float := float64(int.Value)

			return &object.Float{Value: float}
		},
	},
	// floor, then cast to int
	"toint": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			float, ok := args[0].(*object.Float)
			if !ok {
				return newError("argument to `toint` must be a FLOAT, got %s",
					args[0].Type())
			}

			int := int64(float.Value)

			return &object.Integer{Value: int}
		},
	},
}
