
const formInput = document.getElementById('name-0335');
const submitAnker = document.getElementById('submit-anker');

submitAnker.addEventListener('click', function (event) {
    event.preventDefault();
    shortURL(formInput.value);
});




function shortURL(url) {
    const ajaxRequest = {
        url : '/api/v1/url/short',
        type: 'POST',
        dataType : "json",
        data: JSON.stringify({URL: url}),
        beforeSend: function(xhr){xhr.setRequestHeader('Content-Type', 'application/json; charset=utf-8');},
    }

    $.ajax(ajaxRequest).then((res) => {
        console.log('#DEBUG Response done: ', res);

        window.location.replace("/site/result.html?uid=" + res.uid);
    }).catch((err) => {
        console.log('#DEBUG ', err);
        alert("Couldn't make it !")
    });
}

