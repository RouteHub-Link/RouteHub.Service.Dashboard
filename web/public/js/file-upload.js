function TriggerClosestFileInput(el) {
    const fileInput = el.closest('.file-upload').querySelector('input[type="file"]');
    if (fileInput) {
      fileInput.click();
    }
   }
   
   function TriggerShowImage(el) {
     // find closest name="preview" div
     const preview = el.closest('.file-upload').querySelector('[name="preview"]');
     const uploader= el.closest('.file-upload').querySelector('[name="uploader"]');
     uploader.style = "display:none";

     // set style display block
     preview.style.display = "block";
     // find image element under preview div > div > img
     const image = preview.querySelector('img');
   
     // set image src
     image.src = URL.createObjectURL(el.files[0]);

     const fileNameDetailsElement =  preview.childNodes[0].childNodes[0].childNodes[1].childNodes[0];
     const fileSizeDetailsElement =  preview.childNodes[0].childNodes[0].childNodes[1].childNodes[1];
     const fileNameSpan = fileNameDetailsElement.childNodes[0];

     fileNameSpan.textContent= el.files[0].name;
     fileSizeDetailsElement.innerText= `File Size: ${el.files[0].size} bytes`;
   }
   
   function TriggerRemoveInputFile(el) {
     // find  closest name="preview" div
     const preview = el.parentElement.parentElement.parentElement;
     const uploader= preview.parentElement.querySelector('[name="uploader"]');
     uploader.style = "display:flex";

     // set style display none
     preview.style.display = "none";
   
     // find image element under preview div > div > img
     const image = preview.querySelector('img');
   
     // set image src
     image.src = "";
     
   }