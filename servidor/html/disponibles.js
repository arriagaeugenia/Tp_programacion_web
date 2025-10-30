fetch("/obrasDisponibles")
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
                    <hr>
                    `;
                    contenedor.appendChild(obraDiv);
                });
            });