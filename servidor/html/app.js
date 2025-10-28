fetch("/obras")
            .then(res => {
                if (!res.ok) {
                    throw new Error("Error al obtener las obras");
                }
                return res.json();
            })
            .then(data => {
                const contenedor = document.getElementById("contenedor-obras");
                data.forEach(obra => {
                    const obraDiv = document.createElement("div");
                    obraDiv.className = "shadowbox";
                    obraDiv.innerHTML = `
                    <h3>${obra.titulo}</h3>
                    <p><strong>Artista:</strong> ${obra.artista}</p>
                    <p><strong>Descripcion:</strong> ${obra.descripcion}</p>
                    <p><strong>Precio:</strong> ${obra.precio}</p>
                    <p>${obra.vendida}</p>
                    <button type="button" class="eliminar-btn" data-id="${obra.id}">Eliminar</button>
                    <hr>
                    `;
                    contenedor.appendChild(obraDiv);
                });
            });

            document.getElementById("contenedor-obras").addEventListener("click", async (e) => {
            if (e.target.classList.contains("eliminar-btn")) {
                const id = e.target.getAttribute("data-id");
                

                if (confirm("¿Seguro que querés eliminar esta obra?")) {
                    const response = await fetch(`/obras/${id}`, {
                        method: "DELETE",
                        headers: { "Content-Type": "application/json" }
                    });

                    if (response.ok) {
                        alert("Obra eliminada correctamente");
                        e.target.closest(".shadowbox").remove(); // elimina el box del DOM
                    } else {
                        alert("Error al eliminar la obra");
                    }
                }
            }
        });