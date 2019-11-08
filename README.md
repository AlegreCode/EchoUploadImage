# EchoUploadImage
### Plantilla de para subir imágenes al servidor con Echo framework.

##### Packages utilizados:
- **Sprig**: extiende las funcionalidades de template de golang. [Aquí](https://github.com/Masterminds/sprig)
- **gookit/validate**: agrega funciones de validación de entradas. [Aquí](https://github.com/gookit/validate)
- **satori/go.uuid**: agrega implemetación de Universal Unique Identifier(UUID).[Aquí](https://godoc.org/github.com/satori/go.uuid)
- **go-bytesize**: agrega implementación para medir y formatear tamaños de bytes.[Aquí](https://godoc.org/github.com/inhies/go-bytesize#example-ByteSize-Format)


##### Instalación de paquetes
Para evitar conflictos de versiones utilizamos el gestor de dependencias de go [dep](https://golang.github.io/dep/). Debes tener instalado esta herramienta (para ver como clic [aquí](https://golang.github.io/dep/docs/installation.html)), luego entrar a la raíz de tu proyecto y ejecutar el comando `dep ensure`. Una vez instalados todos los paquetes ya correr el proyecto.
