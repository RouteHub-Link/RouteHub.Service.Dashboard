// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Sidebar() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div id=\"hs-application-sidebar\" class=\"hs-overlay [--auto-close:lg] hs-overlay-open:translate-x-0 -translate-x-full transition-all duration-300 transform w-[260px] h-full hidden fixed inset-y-0 start-0 z-[60] bg-white border-e border-gray-200 lg:block lg:translate-x-0 lg:end-auto lg:bottom-0 dark:bg-neutral-800 dark:border-neutral-700\" role=\"dialog\" tabindex=\"-1\" aria-label=\"Sidebar\"><div class=\"relative flex flex-col h-full max-h-full\"><div class=\"px-6 pt-4\"><a class=\"flex-none rounded-xl text-xl inline-block font-semibold focus:outline-none focus:opacity-80 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"#\" aria-label=\"RouteHUB\">RouteHUB</a></div><div class=\"h-full overflow-y-auto [&amp;::-webkit-scrollbar]:w-2 [&amp;::-webkit-scrollbar-thumb]:rounded-full [&amp;::-webkit-scrollbar-track]:bg-gray-100 [&amp;::-webkit-scrollbar-thumb]:bg-gray-300 dark:[&amp;::-webkit-scrollbar-track]:bg-neutral-700 dark:[&amp;::-webkit-scrollbar-thumb]:bg-neutral-500\"><nav class=\"hs-accordion-group p-3 w-full flex flex-col flex-wrap\" data-hs-accordion-always-open><ul class=\"flex flex-col space-y-1\"><li><a class=\"w-full flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"/hubs\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"24px\" viewBox=\"0 -960 960 960\" width=\"24px\" class=\"shrink-0 size-4\" fill=\"currentColor\"><path d=\"M340-320h300q50 0 85-35t35-85q0-50-35-85t-85-35q-8-58-53-99t-101-41q-51 0-92.5 26T332-600q-57 5-94.5 43.5T200-460q0 58 41 99t99 41Zm0-80q-25 0-42.5-17.5T280-460q0-25 17.5-42.5T340-520h60v-20q0-33 23.5-56.5T480-620q33 0 56.5 23.5T560-540v60h80q17 0 28.5 11.5T680-440q0 17-11.5 28.5T640-400H340ZM480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Zm0-320Z\"></path></svg> Hubs</a></li><li><a class=\"w-full flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"/domains\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"24px\" viewBox=\"0 -960 960 960\" width=\"24px\" fill=\"currentColor\"><path d=\"M80-120v-720h400v160h400v560H80Zm80-80h80v-80h-80v80Zm0-160h80v-80h-80v80Zm0-160h80v-80h-80v80Zm0-160h80v-80h-80v80Zm160 480h80v-80h-80v80Zm0-160h80v-80h-80v80Zm0-160h80v-80h-80v80Zm0-160h80v-80h-80v80Zm160 480h320v-400H480v80h80v80h-80v80h80v80h-80v80Zm160-240v-80h80v80h-80Zm0 160v-80h80v80h-80Z\"></path></svg> Domains</a></li><li class=\"hs-accordion\" id=\"settings-accordion\"><a href=\"/settings\" class=\"hs-accordion-toggle w-full text-start flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" aria-expanded=\"true\" aria-controls=\"settings-accordion-child\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"24px\" viewBox=\"0 -960 960 960\" width=\"24px\" fill=\"#e8eaed\"><path d=\"M680-280q25 0 42.5-17.5T740-340q0-25-17.5-42.5T680-400q-25 0-42.5 17.5T620-340q0 25 17.5 42.5T680-280Zm0 120q31 0 57-14.5t42-38.5q-22-13-47-20t-52-7q-27 0-52 7t-47 20q16 24 42 38.5t57 14.5ZM480-80q-139-35-229.5-159.5T160-516v-244l320-120 320 120v227q-19-8-39-14.5t-41-9.5v-147l-240-90-240 90v188q0 47 12.5 94t35 89.5Q310-290 342-254t71 60q11 32 29 61t41 52q-1 0-1.5.5t-1.5.5Zm200 0q-83 0-141.5-58.5T480-280q0-83 58.5-141.5T680-480q83 0 141.5 58.5T880-280q0 83-58.5 141.5T680-80ZM480-494Z\"></path></svg> Settings<!---\n\n\t\t\t\t\t\t\t\t<svg\n\t\t\t\t\t\t\t\t\tclass=\"hs-accordion-active:block ms-auto hidden size-4\"\n\t\t\t\t\t\t\t\t\txmlns=\"http://www.w3.org/2000/svg\"\n\t\t\t\t\t\t\t\t\twidth=\"24\"\n\t\t\t\t\t\t\t\t\theight=\"24\"\n\t\t\t\t\t\t\t\t\tviewBox=\"0 0 24 24\"\n\t\t\t\t\t\t\t\t\tfill=\"none\"\n\t\t\t\t\t\t\t\t\tstroke=\"currentColor\"\n\t\t\t\t\t\t\t\t\tstroke-width=\"2\"\n\t\t\t\t\t\t\t\t\tstroke-linecap=\"round\"\n\t\t\t\t\t\t\t\t\tstroke-linejoin=\"round\"\n\t\t\t\t\t\t\t\t>\n\t\t\t\t\t\t\t\t\t<path d=\"m18 15-6-6-6 6\"></path>\n\t\t\t\t\t\t\t\t</svg>\n\t\t\t\t\t\t\t\t<svg\n\t\t\t\t\t\t\t\t\tclass=\"hs-accordion-active:hidden ms-auto block size-4\"\n\t\t\t\t\t\t\t\t\txmlns=\"http://www.w3.org/2000/svg\"\n\t\t\t\t\t\t\t\t\twidth=\"24\"\n\t\t\t\t\t\t\t\t\theight=\"24\"\n\t\t\t\t\t\t\t\t\tviewBox=\"0 0 24 24\"\n\t\t\t\t\t\t\t\t\tfill=\"none\"\n\t\t\t\t\t\t\t\t\tstroke=\"currentColor\"\n\t\t\t\t\t\t\t\t\tstroke-width=\"2\"\n\t\t\t\t\t\t\t\t\tstroke-linecap=\"round\"\n\t\t\t\t\t\t\t\t\tstroke-linejoin=\"round\"\n\t\t\t\t\t\t\t\t>\n\t\t\t\t\t\t\t\t\t<path d=\"m6 9 6 6 6-6\"></path>\n\t\t\t\t\t\t\t\t</svg>\n\t\t\t\t\t\t\t\t--></a><div id=\"settings-accordion-child\" class=\"hs-accordion-content w-full overflow-hidden transition-[height] duration-300 hidden\" role=\"region\" aria-labelledby=\"settings-accordion\"><ul class=\"ps-8 pt-1 space-y-1\"><li><a class=\"flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"#\">Organization</a></li><li><a class=\"flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"#\">Platform</a></li></ul></div></li><li><a class=\"w-full flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"24px\" viewBox=\"0 -960 960 960\" width=\"24px\" fill=\"currentColor\"><path d=\"M480-80q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80ZM364-182l48-110q-42-15-72.5-46.5T292-412l-110 46q23 64 71 112t111 72Zm-72-366q17-42 47.5-73.5T412-668l-46-110q-64 24-112 72t-72 112l110 46Zm188 188q50 0 85-35t35-85q0-50-35-85t-85-35q-50 0-85 35t-35 85q0 50 35 85t85 35Zm116 178q63-24 110.5-71.5T778-364l-110-48q-15 42-46 72.5T550-292l46 110Zm72-368 110-46q-24-63-71.5-110.5T596-778l-46 112q41 15 71 45.5t47 70.5Z\"></path></svg> Support</a></li></ul></nav></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
