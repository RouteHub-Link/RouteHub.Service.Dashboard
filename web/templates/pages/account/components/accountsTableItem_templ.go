// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func accountsTableItem() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<tr><td class=\"size-px whitespace-nowrap\"></td><td class=\"size-px whitespace-nowrap\"><div class=\"ps-6 lg:ps-3 xl:ps-0 pe-6 py-3\"><div class=\"flex items-center gap-x-3\"><span class=\"inline-flex items-center justify-center size-[38px] max-w-12 max-h-12 h-auto rounded-full bg-white border border-gray-300 dark:bg-neutral-800 dark:border-neutral-700\"><img class=\"inline-block size-[38px] rounded-full max-w-12 max-h-12 h-auto\" src=\"https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&amp;ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&amp;auto=format&amp;fit=facearea&amp;facepad=2&amp;w=320&amp;h=320&amp;q=80\" alt=\"Avatar\"></span><div class=\"grow\"><span class=\"block text-sm font-semibold text-gray-800 dark:text-neutral-200\">Anne Richard</span> <span class=\"block text-sm text-gray-500 dark:text-neutral-500\">anne@site.com</span></div></div></div></td><td class=\"h-px w-72 whitespace-nowrap\"><div class=\"px-6 py-3\"><span class=\"block text-sm font-semibold text-gray-800 dark:text-neutral-200\">Project XYZ</span> <span class=\"block text-sm text-gray-500 dark:text-neutral-500\">Content Creator</span></div><div class=\"px-6 py-3\"><span class=\"block text-sm font-semibold text-gray-800 dark:text-neutral-200\">my-project.link</span> <span class=\"block text-sm text-gray-500 dark:text-neutral-500\">Admin</span></div></td><td class=\"size-px whitespace-nowrap\"><div class=\"px-6 py-3\"><span class=\"py-1 px-1.5 inline-flex items-center gap-x-1 text-xs font-medium bg-teal-100 text-teal-800 rounded-full dark:bg-teal-500/10 dark:text-teal-500\"><svg class=\"size-2.5\" xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"16\" fill=\"currentColor\" viewBox=\"0 0 16 16\"><path d=\"M16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0zm-3.97-3.03a.75.75 0 0 0-1.08.022L7.477 9.417 5.384 7.323a.75.75 0 0 0-1.06 1.06L6.97 11.03a.75.75 0 0 0 1.079-.02l3.992-4.99a.75.75 0 0 0-.01-1.05z\"></path></svg> Active</span></div></td><td class=\"size-px whitespace-nowrap\"><div class=\"px-6 py-3\"><div class=\"flex items-center gap-x-3\"><span class=\"text-xs text-gray-500 dark:text-neutral-500\">5/5</span><div class=\"flex w-full h-1.5 bg-gray-200 rounded-full overflow-hidden dark:bg-neutral-700\"><div class=\"flex flex-col justify-center overflow-hidden bg-gray-800 dark:bg-neutral-200\" role=\"progressbar\" style=\"width: 100%\" aria-valuenow=\"100\" aria-valuemin=\"0\" aria-valuemax=\"100\"></div></div></div></div></td><td class=\"size-px whitespace-nowrap\"><div class=\"px-6 py-3\"><span class=\"text-sm text-gray-500 dark:text-neutral-500\">18 Dec, 15:20</span></div></td><td class=\"size-px whitespace-nowrap\"><div class=\"px-6 py-1.5\"><a class=\"inline-flex items-center gap-x-1 text-sm text-secondary-600 decoration-2 hover:underline focus:outline-none focus:underline font-medium dark:text-amber-600\" href=\"#\"><svg xmlns=\"http://www.w3.org/2000/svg\" height=\"24px\" viewBox=\"0 -960 960 960\" width=\"24px\" fill=\"currentColor\"><path d=\"M200-120q-33 0-56.5-23.5T120-200v-560q0-33 23.5-56.5T200-840h357l-80 80H200v560h560v-278l80-80v358q0 33-23.5 56.5T760-120H200Zm280-360ZM360-360v-170l367-367q12-12 27-18t30-6q16 0 30.5 6t26.5 18l56 57q11 12 17 26.5t6 29.5q0 15-5.5 29.5T897-728L530-360H360Zm481-424-56-56 56 56ZM440-440h56l232-232-28-28-29-28-231 231v57Zm260-260-29-28 29 28 28 28-28-28Z\"></path></svg></a></div></td></tr>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
