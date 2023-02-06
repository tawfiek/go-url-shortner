const params = new Proxy(new URLSearchParams(window.location.search), {
    get: (searchParams, prop) => searchParams.get(prop),
});

let uidValue = params.uid;

const shorterURL = `${window.location.origin}/${uidValue}`;
const resultAnker = document.getElementById('result-anker');

resultAnker.innerText = shorterURL;
resultAnker.setAttribute('href', shorterURL);