# advent2019 - Day 3

## Algorithm

As wires cannot have diagonal segments (either horizontal or vertical), it will be pretty simple to figure out if two segments might cross. This is why I chose to define a wire as a slice of segments and segments as a set of two points with each having x and y coordinates.
The start point for wires will be in coordinates (1;1).