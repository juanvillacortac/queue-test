# Proyecto de colas

Este proyecto simula un sistema de cola de atención, donde se toma en cuenta tres aristas para el ordenamiento de atención, las cuales son la prioridad del cliente basado en su tipo (VIP y regular), la carga de trabajo que se asocia a la atención del cliente, y el tiempo de llegada del cliente de la cola.

Este sistema utiliza un backend realizado con Go donde se implementa como fuente de datos una base de datos SQL (SQLite en este caso), y posee a su vez un monitoreo de la cola en tiempo real usando WebSockets.

El frontend de este sistema está construido utilizando el framework SvelteKit, con un build basado en archivos estáticos, por lo que dichos artefactos se embeben directamente en el servidor de Go y se sirven como un CDN, manejando su respectiva compresión usando gzip y brotli. Al estar integrado el frontend directamente en el servidor de backend, se posee sin ningún costo la posibilidad de usar cookies para autenticación en el mismo dominio.

La autenticación de los usuarios se basa en verificar hashes basados en bcrypt, y se manejan las sesiones sin estado usando JWT, que son enviados al cliente en cookies seguras que no se pueden leer con JavaScript, por lo que se evita su manipulación mediante XSS.

## Diagrama

![Diagrama de arquitectura](docs/Diagrama de colas.png)
