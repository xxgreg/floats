# Floats

This is an experimental Go package to explore the performance potential of using
SIMD assembly to work with slices of floating point numbers.

The API is intended to be compatible with the [Gonum floats package](https://pkg.go.dev/gonum.org/v1/gonum/floats).

It provides AVX2, and Neon assembly implementations for x86, and arm64. It has
been benchmarked on AWS Lambda.

Currently, only float64 is supported.

When running benchmarks on AWS Lambda, performance over plain Go loops is 
improved 2x-4x for x86, and 2x for arm64. The micro-benchmarks haven't been
carefully checked, so it'll be interesting to see how the performance pans
out in real-world code.
