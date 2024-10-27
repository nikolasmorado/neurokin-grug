// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package dashboard

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	t "neurokin/types"
	"neurokin/views/layouts"
)

func Index(u string, t []t.Task) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"min-h-full bg-pink-300 lg:flex lg:flex-1 lg:items-start\"><div class=\"container w-full flex flex-col lg:grid lg:grid-cols-3 items-start justify-start py-10 px-2 lg:py-10 lg:px-8 xl:px-20 lg:gap-x-4 mx-auto\"><section class=\"flex flex-col bg-green-500 p-[1.5px] lg:p-[5px] lg:col-span-1 rounded-[10px] lg:rounded-[16px] w-full\"><div class=\"bg-green-500 flex gap-x-[4px] lg:gap-x-[6px] p-[6px] lg:p-2.5 rounded-t-[10px] lg:rounded-t-[16px]\"><div class=\"bg-white rounded-full w-[9px] h-[9px] lg:w-[12px] lg:h-[12px]\"></div><div class=\"bg-white rounded-full w-[9px] h-[9px] lg:w-[12px] lg:h-[12px]\"></div><div class=\"bg-white rounded-full w-[9px] h-[9px] lg:w-[12px] lg:h-[12px]\"></div></div><div class=\"bg-white rounded-b-[10px] lg:rounded-b-[16px] p-6 \"><div class=\"flex items-start justify-between\"><hgroup class=\"\"><h3 class=\"text-4xl\">Hello,</h3><h2 class=\"text-4xl font-bold\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(u)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/dashboard/index.templ`, Line: 29, Col: 12}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("!</h2></hgroup> <a href=\"/profile/edit\" class=\"border-green-500 border flex justify-between items-center gap-x-3 px-2 py-1 rounded-[4px] bg-pink-300/[0.15]\"><p>Edit</p><svg xmlns=\" http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"#E26FA3\" class=\"size-6\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10\"></path></svg></a></div></div></section><section class=\"bg-white col-span-2 w-full rounded-[16px] border-green-300 border-2 px-4 lg:px-8 py-4 mt-6 lg:mt-0\"><div class=\"w-full grid grid-cols-1 gap-y-4 lg:gap-y-0 lg:grid-cols-2 lg:gap-x-6\"><div class=\"flex flex-col gap-y-6 border border-black rounded-[16px] px-8 py-8\"><h3 class=\"text-2xl font-semibold font-inter\">Intro Video</h3><iframe width=\"374\" height=\"189\" src=\"https://www.youtube-nocookie.com/embed/c4_3S5GOmiY?si=2GkA-TmFH11_upcY\" title=\"YouTube video player\" frameborder=\"0\" allow=\"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share\" referrerpolicy=\"strict-origin-when-cross-origin\" class=\"w-full h-auto min-h-[189px] bg-white border-black rounded-[16px]\" allowfullscreen></iframe></div><div class=\"flex flex-col gap-y-6 border border-black rounded-[16px] px-8 py-8\"><h3 class=\"text-2xl font-semibold\">Test Results</h3><div class=\"min-h-[189px] bg-gray flex items-center justify-center text-center text-black\">[results]</div></div></div><h2 class=\"mt-10 mb-4 font-bold text-3xl\">Tasks</h2><div class=\"grid gap-y-4\"><div class=\"w-full border rounded-[16px] px-8 py-8 grid grid-cols-2\"><div class=\"flex flex-col\"><p>Task</p><h3 class=\"text-2xl font-semibold\">Task Title</h3><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p></div><div class=\"flex flex-col items-end justify-center gap-y-3\"><a href=\"/quiz\" class=\"bg-pink-400 rounded-[4px] border w-[150px] flex items-center justify-center py-3\">Start-> </a><p class=\"text-right opacity-75\">Takes on average <br>25 minutes</p></div></div><div class=\"w-full border rounded-[16px] px-8 py-8 grid grid-cols-2\"><div class=\"flex flex-col\"><p>Task</p><h3 class=\"text-2xl font-semibold\">Task Title</h3><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p></div><div class=\"flex flex-col items-end justify-center gap-y-3\"><a href=\"/quiz\" class=\"bg-pink-400 rounded-[4px] border w-[150px] flex items-center justify-center py-3\">Start-> </a><p class=\"text-right opacity-75\">Takes on average <br>25 minutes</p></div></div></div></section></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
