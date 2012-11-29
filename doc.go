/*
A portable gl layer that aims to abstract and scale. 
Designed for the Go programming language to use with gocos2d. 
*/
package egles

import (
	_ "github.com/mortdeus/egles/khronos/egl"
	_ "github.com/mortdeus/egles/khronos/es/gles2"
	//_ "github.com/mortdeus/egles/gl"
	_ "github.com/mortdeus/egles/khronos/khr"
	_ "github.com/mortdeus/egles/misc"
	//_ "github.com/mortdeus/egles/misc/waffle"
)

/*
TODO:
EGL wrappers 		~30% 		done
GLES2 wrappers 		 ~5% 		done
GLES wrappers 		  0% 		done


Angle,waffle, wayland, nd andriod support have 
yet to be tested but are ideas currently on the road map.  



*/
