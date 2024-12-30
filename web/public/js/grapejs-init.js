function GrapeJSInitializer(id) {
  var editor = grapesjs.init({
    container: `#gjs-${id}`,
    styleManager: {
      clearProperties: 1,
    },
    plugins: [
      "grapesjs-preset-webpage",
      "gjs-blocks-basic",
      BulmaBlocks,
      //'grapesjs-plugin-forms',
      "grapesjs-component-countdown",
      "grapesjs-plugin-export",
      "grapesjs-tabs",
      "grapesjs-custom-code",
      "grapesjs-touch",
      "grapesjs-parser-postcss",
      "grapesjs-tooltip",
      "grapesjs-tui-image-editor",
      "grapesjs-typed",
      "grapesjs-style-bg",
    ],
    canvas: {
      styles: ["https://cdn.jsdelivr.net/npm/bulma@1.0.2/css/bulma.min.css"],
    },
    pluginsOpts: {
      "grapesjs-blocks-basic": {},
      "grapesjs-tabs": {},
      "grapesjs-typed": {},
      "grapesjs-style-bg": {},
      "grapesjs-tooltip": {},
      "grapesjs-tui-image-editor": {},
      "grapesjs-parser-postcss": {},
      "grapesjs-touch": {},
      "grapesjs-custom-code": {},
      "grapesjs-plugin-export": {},
      "grapesjs-component-countdown": {},
      "grapesjs-preset-webpage": {},
    },
    storageManager: {
      type: "local",
      options: {
        local: { key: `gjsProject-${id}` },
      },
    },
  });

  var projectDataEl = document.getElementById(`gjs-data-${id}`);
  var projectHtmlEl = document.getElementById(`gjs-html-${id}`);

  editor.Storage.add("inline", {
    load() {
      return JSON.parse(projectDataEl.value || "{}");
    },
    store(data) {
      const component = editor.Pages.getSelected().getMainComponent();
      projectDataEl.value = JSON.stringify(data);
      projectHtmlEl.value = `<html>
           <head>
             <style>${editor.getCss({ component })}</style>
           </head>
           ${editor.getHtml({ component })}
         <html>`;
    },
  });

  editor.StorageManager.setCurrent("inline");
  editor.StorageManager.load();

  editor.CssComposer.setRule("body", {
    "background-color": "transparent !important;",
  });
}
