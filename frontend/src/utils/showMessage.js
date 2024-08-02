import {Toast} from "bootstrap";

function showSuccessToast(message) {
    const toastContainer = document.createElement('div');
    toastContainer.className = 'toast-container position-fixed top-0 end-0 p-3';
    document.body.appendChild(toastContainer);

    const toastElement = document.createElement('div');
    toastElement.className = 'toast';
    toastElement.innerHTML = `
    <div class="toast-header">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="green" class="bi bi-circle-fill" viewBox="0 0 16 16">
            <circle cx="8" cy="8" r="8"/>
        </svg>
        &ensp;
        <strong class="me-auto">Success</strong>
        <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
    </div>
    <div class="toast-body">
      ${message}
    </div>
  `;
    toastContainer.appendChild(toastElement);

    const toast = new Toast(toastElement);
    toast.show();

}

function showErrorToast(message) {
    const toastContainer = document.createElement('div');
    toastContainer.className = 'toast-container position-fixed top-0 end-0 p-3';
    document.body.appendChild(toastContainer);

    const toastElement = document.createElement('div');
    toastElement.className = 'toast';
    toastElement.innerHTML = `
    <div class="toast-header">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="red" class="bi bi-circle-fill" viewBox="0 0 16 16">
            <circle cx="8" cy="8" r="8"/>
      </svg>
      &ensp;
      <strong class="me-auto">Error</strong>
      <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
    </div>
    <div class="toast-body">
      ${message}
    </div>
  `;
    toastContainer.appendChild(toastElement);

    const toast = new Toast(toastElement);
    toast.show();
}

export {showSuccessToast, showErrorToast};