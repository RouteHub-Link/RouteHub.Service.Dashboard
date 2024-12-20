function TryParseEditorJSData(outputData) {
  var edjsParser = edjsHTML();
  var html = edjsParser.parse(outputData);
  console.log("TryParseEditorJSData;");
  console.log(html);

  return html;
}

function InitializeEditorJS(id, initialData) {
  const editorConfig = {
    readOnly: false,
    holder: id,

    inlineToolbar: ["link", "marker", "bold", "italic"],
    inlineToolbar: true,

    tools: {
      // Each Tool is a Plugin. Pass them via 'class' option with necessary settings {@link docs/tools.md}
      header: {
        class: Header,
        inlineToolbar: ["marker", "link"],
        config: {
          placeholder: "Header",
        },
        shortcut: "CMD+SHIFT+H",
      },

      list: {
        class: NestedList,
        inlineToolbar: true,
        shortcut: "CMD+SHIFT+L",
      },

      checklist: {
        class: Checklist,
        inlineToolbar: true,
      },

      quote: {
        class: Quote,
        inlineToolbar: true,
        config: {
          quotePlaceholder: "Enter a quote",
          captionPlaceholder: "Quote's author",
        },
        shortcut: "CMD+SHIFT+O",
      },

      warning: Warning,

      code: {
        class: CodeTool,
        shortcut: "CMD+SHIFT+C",
      },

      delimiter: Delimiter,

      inlineCode: {
        class: InlineCode,
        shortcut: "CMD+SHIFT+C",
      },

      linkTool: LinkTool,

      raw: RawTool,

      embed: Embed,

      table: {
        class: Table,
        inlineToolbar: true,
        shortcut: "CMD+ALT+T",
      },
    },

    onReady: function () {},
    onChange: function (api, event) {
      console.log("something changed", event);
      var jsonHolder = document.getElementsByName(id + "-json")[0];
      var htmlHolder = document.getElementsByName(id + "-html")[0];
      api.saver.save().then((outputData) => {
        console.log("outputData", outputData);
        jsonHolder.value = JSON.stringify(outputData);
        console.log("html", TryParseEditorJSData(outputData));
        htmlHolder.value = TryParseEditorJSData(outputData);
      });
    },
  };

  if (initialData) {
    parsed = JSON.parse(initialData);
    editorConfig.data = parsed;
  }

  return new EditorJS(editorConfig);
}
