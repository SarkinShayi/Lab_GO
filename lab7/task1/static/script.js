function displayResult(message, isError = false) {
    const resultBlock = document.getElementById("result");
    resultBlock.className = isError ? "result-error" : "result-success";
    resultBlock.innerHTML = message;
}

async function checkResponse(response) {
    if (!response.ok) {
        const errorText = await response.text();
        displayResult(`Помилка: ${errorText}`, true);
        return;
    }
    const data = await response.json();
    displayResult(`Корені: X₁ = ${data.x1}, X₂ = ${data.x2}`);
}

document.getElementById("getForm").addEventListener("submit", function (e) {
    e.preventDefault();

    const a = document.getElementById("getA").value;
    const b = document.getElementById("getB").value;
    const c = document.getElementById("getC").value;

    fetch(`/solve/get?a=${a}&b=${b}&c=${c}`)
        .then(async response => {
            checkResponse(response);
        })
        .catch(error => {
            displayResult(`Помилка при GET-запиті: ${error}`, true);
        });
});

document.getElementById("postForm").addEventListener("submit", function (e) {
    e.preventDefault();

    const a = parseFloat(document.getElementById("postA").value);
    const b = parseFloat(document.getElementById("postB").value);
    const c = parseFloat(document.getElementById("postC").value);

    fetch("/solve/post", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ a, b, c })
    })
        .then(async response => {
            checkResponse(response);
        })
        .catch(error => {
            displayResult(`Помилка при POST-запиті: ${error}`, true);
        });
});
