// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func AccountsTable() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"max-w-[85rem] px-4 py-10 sm:px-6 lg:px-8 lg:py-14 mx-auto\"><div class=\"flex flex-col\"><div class=\"-m-1.5 overflow-x-auto\"><div class=\"p-1.5 min-w-full inline-block align-middle\"><div class=\"bg-white border border-gray-200 rounded-xl shadow-sm overflow-hidden dark:bg-neutral-800 dark:border-neutral-700\"><!-- Header --><div class=\"px-6 py-4 grid gap-3 md:flex md:justify-between md:items-center border-b border-gray-200 dark:border-neutral-700\"><div><h2 class=\"text-xl font-semibold text-gray-800 dark:text-neutral-200\">Accounts</h2><p class=\"text-sm text-gray-600 dark:text-neutral-400\">Invite people, manage teams and more.</p></div><div><div class=\"inline-flex gap-x-2\"><a class=\"py-2 px-3 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-transparent bg-teal-700 text-white hover:bg-blue-700 focus:outline-none focus:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none\" href=\"#\"><svg class=\"shrink-0 size-4\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path> <path d=\"M12 5v14\"></path></svg> Invite</a></div></div></div><!-- End Header --><!-- Table --><table class=\"min-w-full divide-y divide-gray-200 dark:divide-neutral-700\"><thead class=\"bg-gray-50 dark:bg-neutral-800\"><tr><th scope=\"col\" class=\"ps-6 py-3 text-start\"></th><th scope=\"col\" class=\"ps-6 lg:ps-3 xl:ps-0 pe-6 py-3 text-start\"><div class=\"flex items-center gap-x-2\"><span class=\"text-xs font-semibold uppercase tracking-wide text-gray-800 dark:text-neutral-200\">Name</span></div></th><th scope=\"col\" class=\"px-6 py-3 text-start\"><div class=\"flex items-center gap-x-2\"><span class=\"text-xs font-semibold uppercase tracking-wide text-gray-800 dark:text-neutral-200\">Teams</span></div></th><th scope=\"col\" class=\"px-6 py-3 text-start\"><div class=\"flex items-center gap-x-2\"><span class=\"text-xs font-semibold uppercase tracking-wide text-gray-800 dark:text-neutral-200\">Status</span></div></th><th scope=\"col\" class=\"px-6 py-3 text-start\"><div class=\"flex items-center gap-x-2\"><span class=\"text-xs font-semibold uppercase tracking-wide text-gray-800 dark:text-neutral-200\">Links</span></div></th><th scope=\"col\" class=\"px-6 py-3 text-start\"><div class=\"flex items-center gap-x-2\"><span class=\"text-xs font-semibold uppercase tracking-wide text-gray-800 dark:text-neutral-200\">Accepted</span></div></th><th scope=\"col\" class=\"px-6 py-3 text-end\"></th></tr></thead> <tbody class=\"divide-y divide-gray-200 dark:divide-neutral-700\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = accountsTableItem().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</tbody></table><!-- End Table --><!-- Footer --><div class=\"px-6 py-4 grid gap-3 md:flex md:justify-between md:items-center border-t border-gray-200 dark:border-neutral-700\"><div><p class=\"text-sm text-gray-600 dark:text-neutral-400\"><span class=\"font-semibold text-gray-800 dark:text-neutral-200\">12</span> results</p></div><div><div class=\"inline-flex gap-x-2\"><button type=\"button\" class=\"py-1.5 px-2 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none focus:outline-none focus:bg-gray-50 dark:bg-transparent dark:border-neutral-700 dark:text-neutral-300 dark:hover:bg-neutral-800 dark:focus:bg-neutral-800\"><svg class=\"shrink-0 size-4\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m15 18-6-6 6-6\"></path></svg> Prev</button> <button type=\"button\" class=\"py-1.5 px-2 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none focus:outline-none focus:bg-gray-50 dark:bg-transparent dark:border-neutral-700 dark:text-neutral-300 dark:hover:bg-neutral-800 dark:focus:bg-neutral-800\">Next <svg class=\"shrink-0 size-4\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m9 18 6-6-6-6\"></path></svg></button></div></div></div><!-- End Footer --></div></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
