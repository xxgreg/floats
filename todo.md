
* Figure out how to use go generate with this module structure.
* Go module path is wrong due to extra level of nesting. 
* Experiment with unaligned MOV loads to see if the perf improvement is worth it.
* Make a varargs version of AddMany.

Div/DivTo
Sum/Prod

32 bit floats

Clamp/IsFinite

Rate conversions: especially pow(1+r, 1/12) and  pow(1-r, 1/12)

Is using &x[0] to pass *float64 ok? should I be passing the slice triple instead - what does the go runtime need to know?


