
window.asyncer = (function ($) {
  var selectorFormAsync = "form.asyncer";
  var selectorErrorContainer = "#error-container";
  var selectorSuccessContainer = "#success-container";

  // Show a spinner on form submission
  function changeLoadingState(isLoading, form) {
    if (isLoading) {
      var allButtons = document.querySelectorAll("button[type=submit]");
      allButtons.forEach(function (btn) {
        btn.disabled = true;
      });

      form.querySelector("#spinner").classList.remove("d-none");
    } else {
      var allButtons = document.querySelectorAll("button[type=submit]");
      allButtons.forEach(function (btn) {
        btn.disabled = false;
      });
      form.querySelector("#spinner").classList.add("d-none");
    }
  }

  function json(response) {
    return response.json();
  }

  function onFormSubmit(evt) {
    evt.preventDefault();
    var formJQ = $(this);
    formJQ.find(selectorSuccessContainer).hide();
    formJQ.find(selectorErrorContainer).hide();

    var form = $(this).get(0);
    changeLoadingState(true, form);
    var formData = $(this).serialize();
    fetch($(this).attr("action"), {
      method: $(this).attr("method"),
      body: formData,
      headers: {
        "Content-type": "application/x-www-form-urlencoded; charset=UTF-8",
      },
    })
      .then(json)
      .then(function (data) {
        if (data.error !== undefined) {
          formJQ.find(selectorErrorContainer).html(data.error).show();
          changeLoadingState(false, form);
          return;
        }

        if (data.next !== undefined) {
          window.location.href = data.next;
          return;
        }

        if (data.successMsg !== undefined) {
          formJQ.find(selectorSuccessContainer).html(data.successMsg).show();
          changeLoadingState(false, form);

          setTimeout(function () {
            formJQ.find(selectorSuccessContainer).hide();
            if (data.continue !== undefined) {
              window.location.href = data.continue;
            }
          }, 2000);
          return;
        }
      })
      .catch(function (error) {
        console.log(error);
        formJQ.find(selectorErrorContainer).html(error);
        formJQ.find(selectorErrorContainer).show();
        changeLoadingState(false, form);
      });
  }

  function init() {
    $("body").on("submit", selectorFormAsync, onFormSubmit);
  }

  $(function () {
    console.log("start asyncer");
   
    $(".custom-file-input").on("change", function() {
      console.log("clicked");
      var fileName = $(this).val().split("\\").pop();
      $(this).siblings(".custom-file-label").addClass("selected").html(fileName);
    });

    init();
  });

  var asyncer = {};

  return asyncer;
})(jQuery);
