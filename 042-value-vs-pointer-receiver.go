package main

/*
 * Some rules about value vs pointer receivers:
 *
 * A receiver must be a pointer when:
 * - the method needs to modify the receiver
 * - the receiver contains a field that cannot be copied (like type part of sync package)
 *
 * A receiver should be a pointer when it's a large object to prevent costly copying.
 *
 * A receiver must be a value when:
 * - it has to be immutable
 * - if receiver is a map, function or a channel
 *
 * A receiver should be a value when:
 * - it's a slice that doesn't have to be mutated
 * - it's a small array or struct that is a value type
 * - it's a basic type such as int, string or float64
 */
