// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import layouts "RouteHub.Service.Dashboard/web/templates/layouts"
import "github.com/zitadel/oidc/v3/pkg/oidc"

func index(userInfo *oidc.UserInfo) templ.Component {
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
		templ_7745c5c3_Err = layouts.Main(layouts.PageDescription{
			MainContent: homePage(),
			UserInfo:    userInfo,
		}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func homePage() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"container mx-auto\"><div class=\"hs-dropdown [--strategy:absolute] [--flip:false] hs-dropdown-example relative inline-flex\"><button id=\"hs-dropdown-example\" type=\"button\" class=\"hs-dropdown-toggle py-3 px-4 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 focus:outline-none focus:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none dark:bg-neutral-800 dark:border-neutral-700 dark:text-white dark:hover:bg-neutral-700 dark:focus:bg-neutral-700\" aria-haspopup=\"menu\" aria-expanded=\"false\" aria-label=\"Dropdown\">Actions <svg class=\"hs-dropdown-open:rotate-180 size-4 text-gray-600 dark:text-neutral-600\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m6 9 6 6 6-6\"></path></svg></button><div class=\"hs-dropdown-menu transition-[opacity,margin] duration hs-dropdown-open:opacity-100 opacity-0 w-56 hidden z-10 mt-2 min-w-60 bg-white shadow-md rounded-lg p-2 dark:bg-neutral-800 dark:border dark:border-neutral-700 dark:divide-neutral-700\" role=\"menu\" aria-orientation=\"vertical\" aria-labelledby=\"hs-dropdown-example\"><a class=\"flex items-center gap-x-3.5 py-2 px-3 rounded-lg text-sm text-gray-800 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 dark:text-neutral-400 dark:hover:bg-neutral-700 dark:hover:text-neutral-300 dark:focus:bg-neutral-700\" href=\"#\">Newsletter</a> <a class=\"flex items-center gap-x-3.5 py-2 px-3 rounded-lg text-sm text-gray-800 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 dark:text-neutral-400 dark:hover:bg-neutral-700 dark:hover:text-neutral-300 dark:focus:bg-neutral-700\" href=\"#\">Purchases</a> <a class=\"flex items-center gap-x-3.5 py-2 px-3 rounded-lg text-sm text-gray-800 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 dark:text-neutral-400 dark:hover:bg-neutral-700 dark:hover:text-neutral-300 dark:focus:bg-neutral-700\" href=\"#\">Downloads</a> <a class=\"flex items-center gap-x-3.5 py-2 px-3 rounded-lg text-sm text-gray-800 hover:bg-gray-100 focus:outline-none focus:bg-gray-100 dark:text-neutral-400 dark:hover:bg-neutral-700 dark:hover:text-neutral-300 dark:focus:bg-neutral-700\" href=\"#\">Team Account</a></div></div><button type=\"button\" class=\"hs-collapse-toggle py-3 px-4 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-gray-200 bg-white text-gray-800 shadow-sm hover:bg-gray-50 focus:outline-none focus:bg-gray-50 disabled:opacity-50 disabled:pointer-events-none dark:bg-neutral-900 dark:border-neutral-700 dark:text-white dark:hover:bg-neutral-800 dark:focus:bg-neutral-800\" id=\"hs-basic-collapse\" aria-expanded=\"false\" aria-controls=\"hs-basic-collapse-heading\" data-hs-collapse=\"#hs-basic-collapse-heading\">Collapse <svg class=\"hs-collapse-open:rotate-180 size-4\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m6 9 6 6 6-6\"></path></svg></button><div id=\"hs-basic-collapse-heading\" class=\"hs-collapse hidden w-full overflow-hidden transition-[height] duration-300\" aria-labelledby=\"hs-basic-collapse\"><div class=\"mt-5 bg-white rounded-lg py-3 px-4 dark:bg-neutral-800\"><p class=\"text-gray-500 dark:text-neutral-400\">This is a collapse body. It is hidden by default, until the collapse plugin adds the appropriate classes that we use to style each element. These classes control the overall appearance, as well as the showing and hiding via CSS transitions.</p></div></div><div class=\"w-full bg-white rounded-lg shadow-md dark:bg-neutral-800\"><div class=\"hs-accordion-group\"><div class=\"hs-accordion active\" id=\"hs-basic-heading-one\"><button class=\"hs-accordion-toggle hs-accordion-active:text-blue-600 px-6 py-3 inline-flex items-center gap-x-3 text-sm w-full font-semibold text-start text-gray-800 hover:text-gray-500 focus:outline-none focus:text-gray-500 rounded-lg disabled:opacity-50 disabled:pointer-events-none dark:hs-accordion-active:text-blue-500 dark:text-neutral-200 dark:hover:text-neutral-400 dark:focus:text-neutral-400\" aria-expanded=\"true\" aria-controls=\"hs-basic-collapse-one\"><svg class=\"hs-accordion-active:hidden hs-accordion-active:text-blue-600 hs-accordion-active:group-hover:text-blue-600 block size-4 text-gray-600 group-hover:text-gray-500 dark:text-neutral-400\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path> <path d=\"M12 5v14\"></path></svg> <svg class=\"hs-accordion-active:block hs-accordion-active:text-blue-600 hs-accordion-active:group-hover:text-blue-600 hidden size-4 text-gray-600 group-hover:text-gray-500 dark:text-neutral-400\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path></svg> Accordion #1</button><div id=\"hs-basic-collapse-one\" class=\"hs-accordion-content w-full overflow-hidden transition-[height] duration-300\" role=\"region\" aria-labelledby=\"hs-basic-heading-one\"><div class=\"pb-4 px-6\"><p class=\"text-sm text-gray-600 dark:text-neutral-200\">It is hidden by default, until the collapse plugin adds the appropriate classes that we use to style each element.</p></div></div></div><div class=\"hs-accordion\" id=\"hs-basic-heading-two\"><button class=\"hs-accordion-toggle hs-accordion-active:text-blue-600 px-6 py-3 inline-flex items-center gap-x-3 text-sm w-full font-semibold text-start text-gray-800 hover:text-gray-500 focus:outline-none focus:text-gray-500 rounded-lg disabled:opacity-50 disabled:pointer-events-none dark:hs-accordion-active:text-blue-500 dark:text-neutral-200 dark:hover:text-neutral-400 dark:focus:text-neutral-400\" aria-expanded=\"false\" aria-controls=\"hs-basic-collapse-two\"><svg class=\"hs-accordion-active:hidden hs-accordion-active:text-blue-600 hs-accordion-active:group-hover:text-blue-600 block size-4 text-gray-600 group-hover:text-gray-500 dark:text-neutral-400\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path> <path d=\"M12 5v14\"></path></svg> <svg class=\"hs-accordion-active:block hs-accordion-active:text-blue-600 hs-accordion-active:group-hover:text-blue-600 hidden size-4 text-gray-600 group-hover:text-gray-500 dark:text-neutral-400\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path></svg> Accordion #2</button><div id=\"hs-basic-collapse-two\" class=\"hs-accordion-content hidden w-full overflow-hidden transition-[height] duration-300\" role=\"region\" aria-labelledby=\"hs-basic-heading-two\"><div class=\"pb-4 px-6\"><p class=\"text-sm text-gray-600 dark:text-neutral-200\">It is hidden by default, until the collapse plugin adds the appropriate classes that we use to style each element.</p></div></div></div><div class=\"hs-accordion\" id=\"hs-basic-heading-three\"><button class=\"hs-accordion-toggle hs-accordion-active:text-blue-600 px-6 py-3 inline-flex items-center gap-x-3 text-sm w-full font-semibold text-start text-gray-800 hover:text-gray-500 focus:outline-none focus:text-gray-500 rounded-lg disabled:opacity-50 disabled:pointer-events-none dark:hs-accordion-active:text-blue-500 dark:text-neutral-200 dark:hover:text-neutral-400 dark:focus:text-neutral-400\" aria-expanded=\"false\" aria-controls=\"hs-basic-collapse-three\"><svg class=\"hs-accordion-active:hidden hs-accordion-active:text-blue-600 hs-accordion-active:group-hover:text-blue-600 block size-4 text-gray-600 group-hover:text-gray-500 dark:text-neutral-400\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path> <path d=\"M12 5v14\"></path></svg> <svg class=\"hs-accordion-active:block hs-accordion-active:text-blue-600 hs-accordion-active:group-hover:text-blue-600 hidden size-4 text-gray-600 group-hover:text-gray-500 dark:text-neutral-400\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M5 12h14\"></path></svg> Accordion #3</button><div id=\"hs-basic-collapse-three\" class=\"hs-accordion-content hidden w-full overflow-hidden transition-[height] duration-300\" role=\"region\" aria-labelledby=\"hs-basic-heading-three\"><div class=\"pb-4 px-6\"><p class=\"text-sm text-gray-600 dark:text-neutral-200\">It is hidden by default, until the collapse plugin adds the appropriate classes that we use to style each element.</p></div></div></div></div></div><div class=\"w-full bg-white rounded-lg shadow-md dark:bg-neutral-800\"><div class=\"border-b border-gray-200 px-4 dark:border-neutral-700\"><nav class=\"flex gap-x-2\" aria-label=\"Tabs\" role=\"tablist\" aria-orientation=\"horizontal\"><button type=\"button\" class=\"hs-tab-active:font-semibold hs-tab-active:border-blue-600 hs-tab-active:text-blue-600 py-4 px-1 inline-flex items-center gap-x-2 border-b-2 border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50 disabled:pointer-events-none dark:text-neutral-400 dark:hover:text-blue-500 dark:focus:text-blue-500 active\" id=\"basic-tabs-item-1\" aria-selected=\"true\" data-hs-tab=\"#basic-tabs-1\" aria-controls=\"basic-tabs-1\" role=\"tab\">Tab 1</button> <button type=\"button\" class=\"hs-tab-active:font-semibold hs-tab-active:border-blue-600 hs-tab-active:text-blue-600 py-4 px-1 inline-flex items-center gap-x-2 border-b-2 border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50 disabled:pointer-events-none dark:text-neutral-400 dark:hover:text-blue-500 dark:focus:text-blue-500\" id=\"basic-tabs-item-2\" aria-selected=\"false\" data-hs-tab=\"#basic-tabs-2\" aria-controls=\"basic-tabs-2\" role=\"tab\">Tab 2</button> <button type=\"button\" class=\"hs-tab-active:font-semibold hs-tab-active:border-blue-600 hs-tab-active:text-blue-600 py-4 px-1 inline-flex items-center gap-x-2 border-b-2 border-transparent text-sm whitespace-nowrap text-gray-500 hover:text-blue-600 focus:outline-none focus:text-blue-600 disabled:opacity-50 disabled:pointer-events-none dark:text-neutral-400 dark:hover:text-blue-500 dark:focus:text-blue-500\" id=\"basic-tabs-item-3\" aria-selected=\"false\" data-hs-tab=\"#basic-tabs-3\" aria-controls=\"basic-tabs-3\" role=\"tab\">Tab 3</button></nav></div><div class=\"mt-3 p-4\"><div id=\"basic-tabs-1\" role=\"tabpanel\" aria-labelledby=\"basic-tabs-item-1\"><p class=\"text-gray-500 dark:text-neutral-400\">This is the <em class=\"font-semibold text-gray-800 dark:text-neutral-200\">first</em> item's tab body.</p></div><div id=\"basic-tabs-2\" class=\"hidden\" role=\"tabpanel\" aria-labelledby=\"basic-tabs-item-2\"><p class=\"text-gray-500 dark:text-neutral-400\">This is the <em class=\"font-semibold text-gray-800 dark:text-neutral-200\">second</em> item's tab body.</p></div><div id=\"basic-tabs-3\" class=\"hidden\" role=\"tabpanel\" aria-labelledby=\"basic-tabs-item-3\"><p class=\"text-gray-500 dark:text-neutral-400\">This is the <em class=\"font-semibold text-gray-800 dark:text-neutral-200\">third</em> item's tab body.</p></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = stat().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func stat() templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<section class=\"m-10\"><div class=\"relative flex flex-col bg-clip-border rounded-xl bg-white text-gray-700 shadow-md\"><div class=\"p-6 !p-2\"><div class=\"flex gap-2 flex-wrap justify-between px-4 !mt-4\"><h3 class=\"block antialiased tracking-normal font-sans text-3xl font-semibold leading-snug text-blue-gray-900\">$127,092.22</h3><div class=\"flex items-center gap-6\"><div class=\"flex items-center gap-1\"><span class=\"h-2 w-2 bg-blue-500 rounded-full\"></span><p class=\"block antialiased font-sans text-sm font-light leading-normal text-inherit font-normal text-gray-600\">2022</p></div><div class=\"flex items-center gap-1\"><span class=\"h-2 w-2 bg-green-500 rounded-full\"></span><p class=\"block antialiased font-sans text-sm font-light leading-normal text-inherit font-normal text-gray-600\">2023</p></div></div></div><div id=\"chart\"></div></div><div class=\"p-6 flex gap-6 flex-wrap items-center justify-between\"><div><h6 class=\"block antialiased tracking-normal font-sans text-base font-semibold leading-relaxed text-blue-gray-900\">Annual Sales Performance</h6><p class=\"block antialiased font-sans text-sm font-light leading-normal text-inherit font-normal text-gray-600 mt-1\">Year-to-Date sales comparison</p></div><button class=\"align-middle select-none font-sans font-bold text-center uppercase transition-all disabled:opacity-50 disabled:shadow-none disabled:pointer-events-none text-xs py-3 px-6 rounded-lg border border-gray-900 text-gray-900 hover:opacity-75 focus:ring focus:ring-gray-300 active:opacity-[0.85]\" type=\"button\" data-ripple-dark=\"true\">View report</button></div></div></section><script src=\"https://cdn.jsdelivr.net/npm/apexcharts\"></script><script>\n    const options = {\n        series: [\n            {\n                name: \"2022 Sales\",\n                data: [0, 200, 180, 350, 500, 680, 800, 800, 880, 900, 680, 900],\n            },\n            {\n                name: \"2023 Sales\",\n                data: [200, 160, 150, 260, 600, 790, 900, 660, 720, 800, 500, 800],\n            },\n        ],\n        colors: [\"#4CAF50\", \"#2196F3\"],\n        chart: {\n            height: 350,\n            type: \"area\",\n            zoom: {\n                enabled: false,\n            },\n            toolbar: {\n                show: false,\n            },\n        },\n        title: {\n            show: \"\",\n        },\n        dataLabels: {\n            enabled: false,\n        },\n        legend: {\n            show: false,\n        },\n        markers: {\n            size: 0,\n            strokeWidth: 0,\n            strokeColors: \"transparent\",\n        },\n        stroke: {\n            curve: \"smooth\",\n            width: 2,\n        },\n        grid: {\n            show: true,\n            borderColor: \"#EEEEEE\",\n            strokeDashArray: 5,\n            xaxis: {\n                lines: {\n                    show: false,\n                },\n            },\n            padding: {\n                top: 5,\n                right: 20,\n            },\n        },\n        tooltip: {\n            theme: \"light\",\n        },\n        yaxis: {\n            labels: {\n                style: {\n                    colors: \"#757575\",\n                    fontSize: \"12px\",\n                    fontFamily: \"inherit\",\n                    fontWeight: 300,\n                },\n            },\n        },\n        xaxis: {\n            axisTicks: {\n                show: false,\n            },\n            axisBorder: {\n                show: false,\n            },\n            labels: {\n                style: {\n                    colors: \"#757575\",\n                    fontSize: \"12px\",\n                    fontFamily: \"inherit\",\n                    fontWeight: 300,\n                },\n            },\n            categories: [\n                \"Jan\",\n                \"Feb\",\n                \"Mar\",\n                \"Apr\",\n                \"May\",\n                \"Jun\",\n                \"Jul\",\n                \"Aug\",\n                \"Sep\",\n                \"Oct\",\n                \"Nov\",\n                \"Dec\",\n            ],\n        },\n        fill: {\n            type: \"gradient\",\n            gradient: {\n                shadeIntensity: 1,\n                opacityFrom: 0,\n                opacityTo: 0,\n                stops: [0, 100],\n            },\n        },\n    };\n\n    var chart = new ApexCharts(document.querySelector(\"#chart\"), options);\n    chart.render();\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
