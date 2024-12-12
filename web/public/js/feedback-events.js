function tostifyCustomClose(el) {
  const parent = el.closest('.toastify');
  const close = parent.querySelector('.toast-close');

  close.click();
}

document.addEventListener("DOMContentLoaded", function () {

  document.body.addEventListener("close-modal", function (evt) {
    HSOverlay.close(document.getElementById("hs-modal"));
  })


  document.body.addEventListener("toast", function (evt) {

    let messsage = evt.detail.message;

    const toastMarkup = `
<div onclick="tostifyCustomClose(this)" class="hs-removing:translate-x-5 hs-removing:opacity-0 transition duration-300 bg-teal-50 border border-teal-200 text-sm text-teal-800 rounded-lg p-4 dark:bg-teal-800/10 dark:border-teal-900 dark:text-teal-500" role="alert" tabindex="-1" aria-labelledby="hs-dismiss-button-label">
  <div class="flex">
    <div class="shrink-0">
      <svg class="shrink-0 size-4 mt-0.5" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2 2 6.477 2 12s4.477 10 10 10z"></path>
        <path d="m9 12 2 2 4-4"></path>
      </svg>
    </div>
    <div class="ms-2">
      <h3 id="hs-dismiss-button-label" class="text-sm font-medium">
        ${messsage}
      </h3>
    </div>
    <div class="ps-3 ms-auto">
      <div class="-mx-1.5 -my-1.5">
        <button type="button" class="inline-flex bg-teal-50 rounded-lg p-1.5 text-teal-500 hover:bg-teal-100 focus:outline-none focus:bg-teal-100 dark:bg-transparent dark:text-teal-600 dark:hover:bg-teal-800/50 dark:focus:bg-teal-800/50" data-hs-remove-element="#dismiss-alert">
          <span class="sr-only">Dismiss</span>
          <svg class="shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M18 6 6 18"></path>
            <path d="m6 6 12 12"></path>
          </svg>
        </button>
      </div>
    </div>
  </div>
</div>
    `;

    Toastify({
      text: toastMarkup,
      className: "bg-none p-0 hs-toastify-on:opacity-100 opacity-0 fixed -top-[150px] right-[20px] z-[90] transition-all duration-300 w-[320px] text-sm text-gray-700 border border-gray-200 rounded-xl shadow-lg [&>.toast-close]:hidden dark:bg-neutral-800 dark:border-neutral-700 dark:text-neutral-400",
      duration: 3000,
      close: true,
      escapeMarkup: false
    }).showToast();
  })
});

// document.body.dispatchEvent(new Event("toast"));
