const API_BASE = "";

async function loadCourses() {
    const status = document.getElementById("status");
    const grid = document.getElementById("coursesGrid");
    grid.innerHTML = "";
    status.textContent = "Load...";

    try {
        const res = await fetch(`${API_BASE}/course`, {
            headers: { "Accept": "application/json" }
        });

        if (!res.ok) {
            status.textContent = `Ошибка: ${res.status} ${res.statusText}`;
            return;
        }

        const course = await res.json(); // ожидаем массив: [{id, title, ...}]
        if (!Array.isArray(course) || course.length === 0) {
            status.textContent = "Пока нет курсов.";
            return;
        }

        status.textContent = "";

        for (const c of course) {
            const card = document.createElement("a");
            card.className = "card";
            card.href = `/course/${c.id}`; // если страницы курса ещё нет, можно оставить "#"
            card.innerHTML = `
            <div class="card__title">${escapeHtml(c.title ?? "Без названия")}</div>
            <div class="card__meta">ID: ${escapeHtml(String(c.id ?? ""))}</div>
          `;
            grid.appendChild(card);
        }
    } catch (e) {
        status.textContent = "Не удалось загрузить курсы. Проверь backend и CORS.";
    }
}

function escapeHtml(s) {
    return s
        .replaceAll("&", "&amp;")
        .replaceAll("<", "&lt;")
        .replaceAll(">", "&gt;")
        .replaceAll('"', "&quot;")
        .replaceAll("'", "&#039;");
}

document.getElementById("reloadBtn").addEventListener("click", loadCourses);
loadCourses();