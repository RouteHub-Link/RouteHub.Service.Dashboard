// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func HubSidebar() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"hs-application-sidebar-hub\" class=\"hs-overlay [--auto-close:lg] hs-overlay-open:translate-x-0 -translate-x-full transition-all duration-300 transform w-[260px] h-full hidden fixed inset-y-0 start-0 z-[60] bg-white border-e border-gray-200 lg:block lg:translate-x-0 lg:end-auto lg:bottom-0 dark:bg-neutral-800 dark:border-neutral-700\" role=\"dialog\" tabindex=\"-1\" aria-label=\"Sidebar\"><div class=\"relative flex flex-col h-full max-h-full\"><div class=\"px-6 pt-4\"><a class=\"flex-none rounded-xl text-xl inline-block font-semibold focus:outline-none focus:opacity-80 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"#\" aria-label=\"Project\">Project Dashboard <small class=\"text-xs\"><br>My Project</small></a></div><div class=\"h-full overflow-y-auto [&amp;::-webkit-scrollbar]:w-2 [&amp;::-webkit-scrollbar-thumb]:rounded-full [&amp;::-webkit-scrollbar-track]:bg-gray-100 [&amp;::-webkit-scrollbar-thumb]:bg-gray-300 dark:[&amp;::-webkit-scrollbar-track]:bg-neutral-700 dark:[&amp;::-webkit-scrollbar-thumb]:bg-neutral-500\"><nav class=\"hs-accordion-group p-3 w-full flex flex-col flex-wrap\" data-hs-accordion-always-open><ul class=\"flex flex-col space-y-1\"><li><a class=\"flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-500 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"24px\" viewBox=\"0 -960 960 960\" width=\"24px\" fill=\"currentColor\"><path d=\"M640-200 200-480l440-280v560Zm-80-280Zm0 134v-268L350-480l210 134Z\"></path></svg> Main Menu</a></li><li><p class=\"flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-500 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"24px\" viewBox=\"0 -960 960 960\" width=\"24px\" fill=\"currentColor\"><path d=\"M520-520h278q-15-110-91.5-186.5T520-798v278Zm-80 358v-636q-121 15-200.5 105.5T160-480q0 122 79.5 212.5T440-162Zm80 0q110-14 187-91t91-187H520v278Zm-40-318Zm0 400q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 155.5 31.5t127 86q54.5 54.5 86 127T880-480q0 82-31.5 155T763-197.5q-54 54.5-127 86T480-80Z\"></path></svg> Home</p></li><li><a class=\"w-full flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"/hubs\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"24px\" viewBox=\"0 -960 960 960\" width=\"24px\" fill=\"currentColor\"><path d=\"M680-80q-50 0-85-35t-35-85q0-6 3-28L282-392q-16 15-37 23.5t-45 8.5q-50 0-85-35t-35-85q0-50 35-85t85-35q24 0 45 8.5t37 23.5l281-164q-2-7-2.5-13.5T560-760q0-50 35-85t85-35q50 0 85 35t35 85q0 50-35 85t-85 35q-24 0-45-8.5T598-672L317-508q2 7 2.5 13.5t.5 14.5q0 8-.5 14.5T317-452l281 164q16-15 37-23.5t45-8.5q50 0 85 35t35 85q0 50-35 85t-85 35Zm0-80q17 0 28.5-11.5T720-200q0-17-11.5-28.5T680-240q-17 0-28.5 11.5T640-200q0 17 11.5 28.5T680-160ZM200-440q17 0 28.5-11.5T240-480q0-17-11.5-28.5T200-520q-17 0-28.5 11.5T160-480q0 17 11.5 28.5T200-440Zm480-280q17 0 28.5-11.5T720-760q0-17-11.5-28.5T680-800q-17 0-28.5 11.5T640-760q0 17 11.5 28.5T680-720Zm0 520ZM200-480Zm480-280Z\"></path></svg> Links</a></li><li><a class=\"w-full flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"/hubs\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"24px\" viewBox=\"0 -960 960 960\" width=\"24px\" fill=\"currentColor\"><path d=\"M240-120q-45 0-89-22t-71-58q26 0 53-20.5t27-59.5q0-50 35-85t85-35q50 0 85 35t35 85q0 66-47 113t-113 47Zm0-80q33 0 56.5-23.5T320-280q0-17-11.5-28.5T280-320q-17 0-28.5 11.5T240-280q0 23-5.5 42T220-202q5 2 10 2h10Zm230-160L360-470l358-358q11-11 27.5-11.5T774-828l54 54q12 12 12 28t-12 28L470-360Zm-190 80Z\"></path></svg> Customize</a></li><li><a class=\"w-full flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"24px\" viewBox=\"0 -960 960 960\" width=\"24px\" fill=\"currentColor\"><path d=\"M400-480q-66 0-113-47t-47-113q0-66 47-113t113-47q66 0 113 47t47 113q0 66-47 113t-113 47ZM80-160v-112q0-33 17-62t47-44q51-26 115-44t141-18h14q6 0 12 2-8 18-13.5 37.5T404-360h-4q-71 0-127.5 18T180-306q-9 5-14.5 14t-5.5 20v32h252q6 21 16 41.5t22 38.5H80Zm560 40-12-60q-12-5-22.5-10.5T584-204l-58 18-40-68 46-40q-2-14-2-26t2-26l-46-40 40-68 58 18q11-8 21.5-13.5T628-460l12-60h80l12 60q12 5 22.5 11t21.5 15l58-20 40 70-46 40q2 12 2 25t-2 25l46 40-40 68-58-18q-11 8-21.5 13.5T732-180l-12 60h-80Zm40-120q33 0 56.5-23.5T760-320q0-33-23.5-56.5T680-400q-33 0-56.5 23.5T600-320q0 33 23.5 56.5T680-240ZM400-560q33 0 56.5-23.5T480-640q0-33-23.5-56.5T400-720q-33 0-56.5 23.5T320-640q0 33 23.5 56.5T400-560Zm0-80Zm12 400Z\"></path></svg> Members</a></li><li class=\"hs-accordion\" id=\"settings-accordion\"><button type=\"button\" class=\"hs-accordion-toggle w-full text-start flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" aria-expanded=\"true\" aria-controls=\"settings-accordion-child\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"24px\" viewBox=\"0 -960 960 960\" width=\"24px\" fill=\"currentColor\"><path d=\"M680-280q25 0 42.5-17.5T740-340q0-25-17.5-42.5T680-400q-25 0-42.5 17.5T620-340q0 25 17.5 42.5T680-280Zm0 120q31 0 57-14.5t42-38.5q-22-13-47-20t-52-7q-27 0-52 7t-47 20q16 24 42 38.5t57 14.5ZM480-80q-139-35-229.5-159.5T160-516v-244l320-120 320 120v227q-19-8-39-14.5t-41-9.5v-147l-240-90-240 90v188q0 47 12.5 94t35 89.5Q310-290 342-254t71 60q11 32 29 61t41 52q-1 0-1.5.5t-1.5.5Zm200 0q-83 0-141.5-58.5T480-280q0-83 58.5-141.5T680-480q83 0 141.5 58.5T880-280q0 83-58.5 141.5T680-80ZM480-494Z\"></path></svg> Settings <svg class=\"hs-accordion-active:block ms-auto hidden size-4\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m18 15-6-6-6 6\"></path></svg> <svg class=\"hs-accordion-active:hidden ms-auto block size-4\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m6 9 6 6 6-6\"></path></svg></button><div id=\"settings-accordion-child\" class=\"hs-accordion-content w-full overflow-hidden transition-[height] duration-300 hidden\" role=\"region\" aria-labelledby=\"settings-accordion\"><ul class=\"ps-8 pt-1 space-y-1\"><li><a class=\"flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"#\">Domain</a></li><li><a class=\"flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"#\">Deployments</a></li><li><a class=\"flex items-center gap-x-3.5 py-2 px-2.5 text-sm text-gray-800 rounded-lg focus:outline-none focus:bg-zinc-700 hover:bg-zinc-500 dark:hover:bg-neutral-900 dark:text-neutral-200 dark:hover:text-neutral-300\" href=\"#\">Configuration</a></li></ul></div></li></ul></nav></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
