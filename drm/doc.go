package drm

/*
[ModeFbDirtyCmd]
	Mark a region of a framebuffer as dirty.

	Some hardware does not automatically update display contents
	as a hardware or software draw to a framebuffer. This ioctl
	allows userspace to tell the kernel and the hardware what
	regions of the framebuffer have changed.

	The kernel or hardware is free to update more then just the
	region specified by the clip rects. The kernel or hardware
	may also delay and/or coalesce several calls to dirty into a
	single update.

	Userspace may annotate the updates, the annotates are a
	promise made by the caller that the change is either a copy
	of pixels or a fill of a single color in the region specified.

	If the DRM_MODE_FB_DIRTY_ANNOTATE_COPY flag is given then
	the number of updated regions are half of num_clips given,
	where the clip rects are paired in src and dst. The width and
	height of each one of the pairs must match.

	If the DRM_MODE_FB_DIRTY_ANNOTATE_FILL flag is given the caller
	promises that the region specified of the clip rects is filled
	completely with a single color as given in the color argument.

[ModeFbCmd2]	
	In case of planar formats, this ioctl allows up to 4
	buffer objects with offsets and pitches per plane.
	The pitch and offset order is dictated by the fourcc,
	e.g. NV12 (http://fourcc.org/yuv.php#NV12) is described as:

	YUV 4:2:0 image with a plane of 8 bit Y samples
	followed by an interleaved U/V plane containing
	8 bit 2x2 subsampled colour difference samples.

	so it would consist of Y as offset[0] and UV as
	offset[1].  Note that offset[0] will generally
	be 0.
[ModeCtrcPageFlip]
	Request a page flip on the specified crtc.

	This ioctl will ask KMS to schedule a page flip for the specified
	crtc.  Once any pending rendering targeting the specified fb (as of
	ioctl time) has completed, the crtc will be reprogrammed to display
	that fb after the next vertical refresh.  The ioctl returns
	immediately, but subsequent rendering to the current fb will block
	in the execbuffer ioctl until the page flip happens.  If a page
	flip is already pending as the ioctl is called, EBUSY will be
	returned.

	The ioctl supports one flag, DRM_MODEPAGE_FLIPEVENT, which will
	request that drm sends back a vblank event (see drm.h: struct
	event_vblank) when the page flip is done.  The user_data field
	passed in with this ioctl will be returned as the user_data field
	in the vblank event struct.

	The reserved field must be zero until we figure out something
	clever to use it for.
*/

/*	Copyright (c) 1999 Precision Insight, Inc., Cedar Park, Texas.
	Copyright (c) 2000 VA Linux Systems, Inc., Sunnyvale, California.
 	Copyright (c) 2007 Dave Airlie <airlied@linux.ie>
 	Copyright (c) 2007 Jakob Bornecrantz <wallbraker@gmail.com>
 	Copyright (c) 2008 Red Hat Inc.
 	Copyright (c) 2007-2008 Tungsten Graphics, Inc., Cedar Park, TX., USA
 	Copyright (c) 2007-2008 Intel Corporation
 	Copyright (c) 2012 mortdeus <mortdeus@gocos2d.org>

 	Permission is hereby granted, free of charge, to any person obtaining a
 	copy of this software and associated documentation files (the "Software"),
 	to deal in the Software without restriction, including without limitation
 	the rights to use, copy, modify, merge, publish, distribute, sublicense,
 	and/or sell copies of the Software, and to permit persons to whom the
 	Software is furnished to do so, subject to the following conditions:

 	The above copyright notice and this permission notice shall be included in
 	all copies or substantial portions of the Software.

 	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 	FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
 	IN THE SOFTWARE.*/
var License int
