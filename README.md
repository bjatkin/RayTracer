# GO Ray / Path Tracer

This is a small raytracer writen in go for CS 655 Advanced Compouter Graphics.
It supports several primative shapes such as triangles and spheres as well as several
types of lights (area light, point light). It uses oct-tree splitting and bounding boxes
to speed up the render times. It also makes use of go-routines to calculate multiples
rays at a time.

# How To Use

You can complie this project by running the build.sh script. This will produce a binary called tracer.
you can run tracer 2 optional arguments. ```tracer [ray|path|both] [output file name]``` ray with run
the ray tracer, path will run the path tracer, both will run both the ray tracer and path tracer. The
output file is the location where the output will be saved. Ray traced images will be prefaced with RT_,
Path traced files will be prefaced with PT_.

# Path Traced Examples

![PT Chess](https://github.com/bjatkin/RayTracer/blob/master/Renders/PathTracedChessScean.png)
![PT Cornel](https://github.com/bjatkin/RayTracer/blob/master/Renders/PathTracedCornelBox.png)
![PT Reflect](https://github.com/bjatkin/RayTracer/blob/master/Renders/PathTracedReflectScean.png)

# Ray Traced Examples

![RT Chess](https://github.com/bjatkin/RayTracer/blob/master/Renders/RayTracedChessScean.png)
![RT Cornel](https://github.com/bjatkin/RayTracer/blob/master/Renders/RayTracedCornelBox.png)
![RT Cornel 2](https://github.com/bjatkin/RayTracer/blob/master/Renders/RayTracedCornelBox2.png)
![RT Chess 2](https://github.com/bjatkin/RayTracer/blob/master/Renders/RayTracedScean.png)

# TODO
 * Add in ability to import obj files
 * Add in texture maping
 * Clean up the code
    * make file names uniform
    * make switching back and forth between ray/path tracing easier
    * configure multi-threading stuff easier
    * configure other options like rays per pixel/ bounces etc. easier
 * Add more example renders
 * Parse data from a file instead of hard coding the scean