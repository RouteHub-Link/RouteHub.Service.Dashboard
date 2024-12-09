class t extends HTMLElement{get src(){return this.getAttribute("src")}set src(t){t?this.setAttribute("src",t):this.removeAttribute("src")}get auto(){return!this.hasAttribute("no-auto")}get bodyClass(){const t=this.getAttribute("body-class");return"markdown-body"+(t?" "+t:"")}constructor(){super();try{this.version="3.1.6"}catch{}this.template="";this._clicked=(t=>{var e;if(t.metaKey||t.ctrlKey||t.altKey||t.shiftKey||t.defaultPrevented)return;const s=null==(e=t.target)?void 0:e.closest("a");s&&s.hash&&s.host===location.host&&s.pathname===location.pathname&&this.goto(s.hash)}).bind(this),this._observer=new MutationObserver((()=>{this._observe(),this.auto&&this.render()})),this._loaded=!1,this.root=this}static get observedAttributes(){return["src","body-class"]}attributeChangedCallback(t,e,s){var a;if(this.ready&&e!==s)switch(t){case"body-class":null==(a=this.root.querySelector(".markdown-body"))||a.setAttribute("class",this.bodyClass);break;case"src":this.auto&&this.render()}}async connectedCallback(){var t;this._loaded||(await this.load(),this.hasAttribute("no-shadow")||(this.root=this.attachShadow({mode:"open"})),this.root.prepend(this.frag(`<div class="markdown-styles"></div><div class="${this.bodyClass}"></div>`)),this._loaded=!0),null==(t=this.shadowRoot)||t.addEventListener("click",this._clicked),this._observer.observe(this,{childList:!0}),this._observe(),this.ready=!0,this.fire("zero-md-ready"),this.auto&&this.render()}disconnectedCallback(){var t;null==(t=this.shadowRoot)||t.removeEventListener("click",this._clicked),this._observer.disconnect(),this.ready=!1}_observe(){this.querySelectorAll('template,script[type="text/markdown"]').forEach((t=>this._observer.observe(t.content||t,{childList:!0,subtree:!0,attributes:!0,characterData:!0})))}async load(){}async parse({text:t=""}){return t}goto(t){var e;const s=this.shadowRoot||document;t&&(null==(e=s.getElementById(decodeURIComponent("#"===t[0]?t.slice(1):t)))||e.scrollIntoView())}frag(t){const e=document.createElement("template");return e.innerHTML=t,e.content}hash(t){let e=5381;for(let s=0;s<t.length;s++)e=(e<<5)+e^t.charCodeAt(s);return(e>>>0).toString(36)}tick(){return new Promise((t=>requestAnimationFrame(t)))}fire(t,e={}){this.dispatchEvent(new CustomEvent(t,{detail:e,bubbles:!0}))}async read(t){const{target:e}=t,s=(s="",a="")=>{var r;const i=this.hash(s),n=(null==(r=this.root.querySelector(`.markdown-${e}`))?void 0:r.getAttribute("data-hash"))!==i;return{...t,text:s,hash:i,changed:n,baseUrl:a}};switch(e){case"styles":{const t=(t="")=>{var e;return null==(e=this.querySelector(t))?void 0:e.innerHTML};return s((t("template[data-prepend]")??"")+(t("template:not([data-prepend],[data-append])")??this.template)+(t("template[data-append]")??""))}case"body":{if(this.src){const t=await fetch(this.src);if(t.ok){const e=()=>{const t=document.createElement("a");return t.href=this.src||"",t.href.substring(0,t.href.lastIndexOf("/")+1)};return s(await t.text(),e())}console.warn("[zero-md] error reading src",this.src)}const t=this.querySelector('script[type="text/markdown"]');return s((null==t?void 0:t.text)||"")}default:return s()}}async stamp(t){const{target:e,text:s="",hash:a=""}=t,r=this.root.querySelector(`.markdown-${e}`);if(!r)return t;r.setAttribute("data-hash",a);const i=this.frag(s),n=Array.from(i.querySelectorAll('link[rel="stylesheet"]')||[]),o=Promise.all(n.map((t=>new Promise((e=>{t.onload=e,t.onerror=s=>{console.warn("[zero-md] error loading stylesheet",t.href),e(s)}})))));return r.innerHTML="",r.append(i),await o,{...t,stamped:!0}}async render({fire:t=!0,goto:e=location.hash}={}){const s=await this.read({target:"styles"}),a=s.changed&&this.stamp(s),r=await this.read({target:"body"});if(r.changed){const t=this.parse(r);await a,await this.tick(),await this.stamp({...r,text:await t})}else await a;await this.tick();const i={styles:s.changed,body:r.changed};return t&&this.fire("zero-md-rendered",i),this.auto&&e&&this.goto(e),i}}const e=/^(\${1,2})(?!\$)((?:\\.|[^\\\n])*?(?:\\.|[^\\\n\$]))\1(?=[\s?!\.,:？！。，：]|$)/,s=/^(\${1,2})(?!\$)((?:\\.|[^\\\n])*?(?:\\.|[^\\\n\$]))\1/,a=/^(\${1,2})\n((?:\\[^]|[^\\])+?)\n\1(?:\n|$)/;function r(t={}){return{extensions:[i(t,(t=>t.text)),n(t,(t=>t.text))]}}function i(t,a){const r=t&&t.nonStandard,i=r?s:e;return{name:"inlineKatex",level:"inline",start(t){let e,s=t;for(;s;){if(e=s.indexOf("$"),-1===e)return;if(r?e>-1:0===e||" "===s.charAt(e-1)){if(s.substring(e).match(i))return e}s=s.substring(e+1).replace(/^\$+/,"")}},tokenizer(t,e){const s=t.match(i);if(s)return{type:"inlineKatex",raw:s[0],text:s[2].trim(),displayMode:2===s[1].length}},renderer:a}}function n(t,e){return{name:"blockKatex",level:"block",tokenizer(t,e){const s=t.match(a);if(s)return{type:"blockKatex",raw:s[0],text:s[2].trim(),displayMode:2===s[1].length}},renderer:e}}const o=t=>`https://cdn.jsdelivr.net/npm/${t}`,d=(t,e)=>`<link rel="stylesheet" href="${t}"${e?` ${e}`:""}>`,h=async(t,e="default")=>(await import(t))[e],l={HOST:"<style>:host{display:block;position:relative;contain:content;}:host([hidden]){display:none;}</style>",MARKDOWN:d(o("github-markdown-css@5/github-markdown.min.css")),MARKDOWN_LIGHT:d(o("github-markdown-css@5/github-markdown-light.min.css")),MARKDOWN_DARK:d(o("github-markdown-css@5/github-markdown-dark.min.css")),HIGHLIGHT_LIGHT:d(o("@highlightjs/cdn-assets@11/styles/github.min.css")),HIGHLIGHT_DARK:d(o("@highlightjs/cdn-assets@11/styles/github-dark.min.css")),HIGHLIGHT_PREFERS_DARK:d(o("@highlightjs/cdn-assets@11/styles/github-dark.min.css"),'media="(prefers-color-scheme:dark)"'),KATEX:d(o("katex@0/dist/katex.min.css")),preset(t=""){const{HOST:e,MARKDOWN:s,MARKDOWN_LIGHT:a,MARKDOWN_DARK:r,HIGHLIGHT_LIGHT:i,HIGHLIGHT_DARK:n,HIGHLIGHT_PREFERS_DARK:o,KATEX:d}=this,h=t=>`${e}${t}${d}`;switch(t){case"light":return h(a+i);case"dark":return h(r+n);default:return h(s+i+o)}}},c={marked:async()=>new(await h(o("marked@14/lib/marked.esm.js"),"Marked"))({async:!0}),markedBaseUrl:()=>h(o("marked-base-url@1/+esm"),"baseUrl"),markedHighlight:()=>h(o("marked-highlight@2/+esm"),"markedHighlight"),markedGfmHeadingId:()=>h(o("marked-gfm-heading-id@4/+esm"),"gfmHeadingId"),markedAlert:()=>h(o("marked-alert@2/+esm")),hljs:()=>h(o("@highlightjs/cdn-assets@11/es/highlight.min.js")),mermaid:()=>h(o("mermaid@11/dist/mermaid.esm.min.mjs")),katex:()=>h(o("katex@0/dist/katex.mjs"))};let m,u,g,k=0;class y extends t{async load(t={}){const{marked:e,markedBaseUrl:s,markedHighlight:a,markedGfmHeadingId:i,markedAlert:n,hljs:o,mermaid:d,katex:h,katexOptions:y={nonStandard:!0,throwOnError:!1}}={...c,...t};this.template=l.preset();const b=await Promise.all([e(),s(),i(),n(),a()]);this.marked=b[0],this.setBaseUrl=b[1];const p=async(t,e)=>(g||(g=await h()),`${g.renderToString(t,{...y,displayMode:e})}${e?"":"\n"}`);this.marked.use(b[2](),b[3](),{...b[4]({async:!0,highlight:async(t,e)=>{if("mermaid"===e){u||(u=await d(),u.initialize({startOnLoad:!1}));const{svg:e}=await u.render("mermaid-svg-"+k++,t);return e}return"math"===e?await p(t,!0):(m||(m=await o()),m.getLanguage(e)?m.highlight(t,{language:e}).value:m.highlightAuto(t).value)}}),renderer:{code:({text:t,lang:e})=>"mermaid"===e?`<div class="mermaid">${t}</div>`:"math"===e?t:`<pre><code class="hljs${e?` language-${e}`:""}">${t}\n</code></pre>`}},{...r(y),walkTokens:async t=>{["inlineKatex","blockKatex"].includes(t.type)&&(t.text=await p(t.text,"blockKatex"===t.type))}})}async parse({text:t,baseUrl:e}){return this.marked.use(this.setBaseUrl(e||"")),this.marked.parse(t)}}new URL(import.meta.url).searchParams.has("register")&&customElements.define("zero-md",y);export{c as LOADERS,l as STYLES,t as ZeroMdBase,y as default};