window.addEventListener("load", () => {
    fetch("/api/init")
        .then(res => res.json())
        .then(data => {
            document.getElementById("array").innerHTML = `<strong>Сгенерований зріз:</strong> [${data.map(n => n.toFixed(1)).join(", ")}]`;
        })
        .catch(err => {
            document.getElementById("array").innerHTML = `<span style="color:red;">Помилка при генерації: ${err}</span>`;
        });
});

function processData(method) {
    const url = method === "get" ? "/api/get" : "/api/post";
    const options = method === "post" ? { method: "POST" } : {};

    fetch(url, options)
        .then(res => res.json())
        .then(data => {
            const html = `
                <p><strong>Сума від'ємних елементів:</strong> ${data.neg_sum.toFixed(2)}</p>
                <p><strong>Добуток всіх елементів:</strong> ${data.multiplying.toFixed(2)}</p>
            `;
            document.getElementById("result").innerHTML = html;
        })
        .catch(err => {
            document.getElementById("result").innerHTML = `<span style="color:red;">Помилка: ${err}</span>`;
        });
}
