// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package domain

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func CreateDomain(messageTemplate templ.Component, renderForm bool) templ.Component {
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
		if messageTemplate != nil {
			templ_7745c5c3_Err = messageTemplate.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if renderForm {
			templ_7745c5c3_Err = form().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func form() templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"create_domain_form_center\"><form hx-post=\"/domains/create\" hx-indicator=\"#spinner\" hx-target=\"#create_domain_form_center\" hx-swap=\"innerHTML\"><div class=\"grid gap-y-4\"><!-- Form Group --><div class=\"max-w-sm space-y-3\"><div><label for=\"hs-inline-add-on\" class=\"block text-sm font-medium mb-2 dark:text-white\">Website URL</label><div class=\"relative\"><input type=\"text\" id=\"hs-inline-add-on\" name=\"hs-inline-add-on\" class=\"py-3 px-4 ps-16 block w-full border-gray-200 shadow-sm rounded-lg text-sm focus:z-10 focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none dark:bg-neutral-900 dark:border-neutral-700 dark:text-neutral-400 dark:placeholder-neutral-500 dark:focus:ring-neutral-600\" placeholder=\"www.example.com\"><div class=\"absolute inset-y-0 start-0 flex items-center pointer-events-none z-20 ps-4\"><span class=\"text-sm text-gray-500 dark:text-neutral-500\">http://</span></div></div></div></div><div><label for=\"name\" class=\"block text-sm mb-2 dark:text-white\">Domain Name</label><div class=\"relative\"><input type=\"text\" name=\"name\" class=\"py-3 px-4 block w-full border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none dark:bg-neutral-900 dark:border-neutral-800 dark:text-neutral-400 dark:placeholder-neutral-500 dark:focus:ring-neutral-600\" required placeholder=\"example.com\" aria-describedby=\"name-error\"></div><p class=\"hidden text-xs text-red-600 mt-2\" id=\"name-error\">Please include a valid Domain name.</p></div><div><label for=\"Domain\" class=\"block text-sm mb-2 dark:text-white\">Domain address</label><div class=\"relative\"><input type=\"url\" name=\"url\" class=\"py-3 px-4 block w-full border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none dark:bg-neutral-900 dark:border-neutral-800 dark:text-neutral-400 dark:placeholder-neutral-500 dark:focus:ring-neutral-600\" required placeholder=\"https://\" aria-describedby=\"domain-error\"></div><p class=\"hidden text-xs text-red-600 mt-2\" id=\"domain-error\">Please include a valid Domain address so we can use for deployment & proxy settings.</p></div><button type=\"submit\" class=\"relative py-3 px-4 inline-flex justify-center items-center gap-2 text-sm font-medium rounded-lg border border-transparent bg-blue-600 text-white hover:bg-blue-700 focus:outline-none focus:bg-blue-700 disabled:opacity-50 disabled:pointer-events-none\">Add Domain<div class=\"htmx-indicator absolute w-full h-full pt-3 dark:bg-gray-800/[.4]\" id=\"spinner\"><span class=\"animate-spin inline-block w-4 h-4 border-[3px] border-current border-t-transparent text-white rounded-full\" role=\"status\" aria-label=\"loading\" id=\"spinner\"></span></div></button></div></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate