function QuillInitializer(id) {
  var valueInput = document.getElementsByName(id)[0];
  var quill = new Quill("#" + id, {
    modules: {
      toolbar: [
        ["bold", "italic", "underline", "strike"], // toggled buttons
        ["blockquote", "code-block"],
        ["link", "image", "video", "formula"],

        [{ header: 1 }, { header: 2 }], // custom button values
        [{ list: "ordered" }, { list: "bullet" }, { list: "check" }],
        [{ script: "sub" }, { script: "super" }], // superscript/subscript
        [{ indent: "-1" }, { indent: "+1" }], // outdent/indent
        [{ direction: "rtl" }], // text direction

        [{ size: ["small", false, "large", "huge"] }], // custom dropdown
        [{ header: [1, 2, 3, 4, 5, 6, false] }],

        [{ color: [] }, { background: [] }], // dropdown with defaults from theme
        [{ font: [] }],
        [{ align: [] }],
      ],
    },
    placeholder: "Compose an epic...",
    theme: "snow",
  });

  function update() {
    var html = quill.root.innerHTML;
    valueInput.value = html;
  }

  quill.on(Quill.events.TEXT_CHANGE, update);
}
