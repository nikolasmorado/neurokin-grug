// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package quiz

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "strconv"

func Qbox(q, l, u, s, i string, a []string, t string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form hx-put=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs("/quiz/" + s)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/quiz/qbox.templ`, Line: 7, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"submit\" hx-target=\"#content-box\" hx-swap=\"outerHTML\" class=\"bg-white rounded-b-[10px] lg:rounded-b-[16px] p-16\" id=\"content-box\"><p>Question ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(l)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/quiz/qbox.templ`, Line: 14, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("/")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(u)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/quiz/qbox.templ`, Line: 14, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"flex w-full justify-between items-center pt-2 pb-6\"><h1 class=\"font-semibold text-2xl\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(q)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/quiz/qbox.templ`, Line: 17, Col: 7}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><button type=\"button\" id=\"bookmark\" class=\"min-w-[24px] p-4 min-h-[24px] flex items-center justify-center z-10\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" stroke-width=\"1.5\" stroke=\"currentColor\" class=\"size-6 min-w-[24px] min-h-[24px] cursor-pointer z-[-1]\" id=\"bookmark-svg\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0 1 11.186 0Z\"></path></svg></button></div><div class=\"flex flex-col gap-y-3 pb-7\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if t == "mc" {
			for i, t := range a {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button type=\"button\" data-index=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(i + 1))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/quiz/qbox.templ`, Line: 43, Col: 38}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"bg-pink-300 flex px-2 py-[7px] items-center border rounded-[4px] gap-x-2 font-inter w-fit hover:bg-pink-500 transition-all option-btn\"><span class=\"bg-white border rounded-full w-[20px] h-[20px] flex items-center justify-center\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(i + 1))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/quiz/qbox.templ`, Line: 47, Col: 28}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span><p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(t)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/quiz/qbox.templ`, Line: 49, Col: 12}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <input type=\"hidden\" name=\"selectedOption\" id=\"selectedOption\"> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else if t == "text" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <textarea name=\"textAnswer\" id=\"textAnswer\" class=\"border rounded-[4px] p-2 font-inter w-full\" placeholder=\"Write your answer here...\"></textarea> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else if t == "dropdown" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <select name=\"selectedOption\" id=\"dropdownSelect\" class=\"border rounded-[4px] p-2 font-inter w-full\"><option value=\"\">Select an option...</option> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for i, t := range a {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var9 string
				templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(i + 1))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/quiz/qbox.templ`, Line: 70, Col: 41}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var10 string
				templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(t)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/quiz/qbox.templ`, Line: 70, Col: 47}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"hidden\" name=\"bookmarked\" id=\"bookmarked\" value=\"false\"> <input type=\"hidden\" name=\"questionId\" id=\"questionId\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(i)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/quiz/qbox.templ`, Line: 75, Col: 67}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div id=\"contextBox\" class=\"hidden flex flex-col gap-y-2\"><textarea name=\"additionalContext\" class=\"border rounded-[4px] p-2 font-inter\" placeholder=\"Add additional context here...\"></textarea></div><div class=\"flex gap-x-2 items-center\"><button type=\"button\" id=\"addContextBtn\" class=\"bg-green-300 flex px-4 py-[4px] items-center border rounded-[4px] font-inter w-fit\"><p>Add additional context</p></button><p class=\"text-xs text-gray-400 text-[12px] opacity-50\">OR</p><button type=\"submit\" id=\"skipBtn\" class=\"bg-green-700 flex px-4 py-[4px] items-center border rounded-[4px] font-inter w-fit\"><p>Skip</p></button></div></div><a href=\"#\" class=\"text-black text-[12px] opacity-75 mt-8 font-inter\">&lt;- Previous Question</a></form><script>\n  function attachEventListeners() {\n    const contentBox = document.getElementById('content-box');\n    const skipBtn = document.getElementById('skipBtn');\n    const textAnswer = document.getElementById('textAnswer');\n    const dropdownSelect = document.getElementById('dropdownSelect');\n\n    contentBox.querySelectorAll('.option-btn').forEach(button => {\n      button.addEventListener('click', function (e) {\n        e.preventDefault();\n        selectOption(e.target.closest('.option-btn'));\n      });\n    });\n\n    const ctxBtn = document.getElementById('addContextBtn');\n    ctxBtn.addEventListener('click', function (e) {\n      e.preventDefault();\n      toggleContextBox();\n    });\n\n    const bookmarkBtn = document.getElementById('bookmark');\n    bookmarkBtn.addEventListener('click', function (e) {\n      e.preventDefault();\n      toggleBookmark();\n    });\n\n    if (textAnswer) {\n      textAnswer.addEventListener('input', function () {\n        skipBtn.innerText = textAnswer.value.trim() !== \"\" ? \"Next\" : \"Skip\";\n      });\n    }\n\n    if (dropdownSelect) {\n      dropdownSelect.addEventListener('change', function () {\n        skipBtn.innerText = dropdownSelect.value !== \"\" ? \"Next\" : \"Skip\";\n      });\n    }\n\n    document.addEventListener('keydown', function (e) {\n      if (e.key >= '1' && e.key <= '4') {\n        e.preventDefault();\n        const optionIndex = parseInt(e.key, 10) - 1;\n        const optionBtn = contentBox.querySelectorAll('.option-btn')[optionIndex];\n        if (optionBtn) {\n          selectOption(optionBtn);\n        }\n      } else if (e.key === 'Enter') {\n        e.preventDefault();\n        if ((textAnswer && textAnswer.value.trim() !== \"\") || (dropdownSelect && dropdownSelect.value !== \"\")) {\n          skipBtn.click();\n        } else {\n          contentBox.dispatchEvent(new Event('submit', { bubbles: true, cancelable: true }));\n        }\n      }\n    });\n  }\n\n  function selectOption(selectedOption) {\n    const contentBox = document.getElementById('content-box');\n    contentBox.querySelectorAll('.option-btn').forEach(btn => btn.classList.remove('bg-pink-500'));\n    selectedOption.classList.add('bg-pink-500');\n    document.getElementById('selectedOption').value = selectedOption.getAttribute('data-index');\n    document.getElementById('skipBtn').innerText = \"Next\";\n  }\n\n  function toggleContextBox() {\n    const contextBox = document.getElementById('contextBox');\n    const ctxBtn = document.getElementById('addContextBtn');\n    ctxBtn.innerText = contextBox.classList.toggle('hidden') ? \"Add additional context\" : \"Hide\";\n  }\n\n  function toggleBookmark() {\n    const bookmark = document.getElementById('bookmark-svg');\n    const bookmarkState = document.getElementById('bookmarked');\n    const filled = bookmark.getAttribute('fill') === 'none';\n    bookmark.setAttribute('fill', filled ? '#F5A6D5' : 'none');\n    bookmarkState.value = filled ? \"true\" : \"false\";\n  }\n\n  document.addEventListener('htmx:afterSwap', function (event) {\n    if (event.detail.target.id === 'content-box') {\n      attachEventListeners();\n    }\n  });\n\n  attachEventListeners();\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
