// Peticfiones AJAX
// son peticiones que se realizan por medio de javaScript y son peticiones asincronicas
// su nombre es AJAX, por que antes la X final significaba XML., despues nacio JSON, que es mas liviano que XML
// AJAX, fu creado por microsoft e Internet Explorer
// JQUERY y ANGULAR son algunas dde las herramientas que ya cuentan con duncionalidades propias para usar AJAX
function peticionAjax(metodo, url, obj) {
  return new Promise(function(resolver, rechazar) {
    // para realizar peticiones de AJAX se necesita de un objeto de javaScript llamado XMLHttpRequest, este es el cargado de hacer las peticiones y recibir la respuesta ddel servidor
    // let es la forma de declarar una variable en java script con un ambito local, a partir de la version enmascript6
    let xhr = new XMLHttpRequest();
    // se abre la conexion
    // Metodos: post(crear), get(obtener, consultar), delete(borrar), put(actualizar)
    // el tercer metodo pregunta si la peticion es asicrona o no, se coloca true para decirle que si
    xhr.open(metodo, url, true);
    // se coloca el tipotenido a enviar
    xhr.setRequestHeader('Content-Type', 'application/json');
    if (sessionStorage.getItem('token')) {
    xhr.setRequestHeader('Authorization', sessionStorage.getItem('token'))
  }
    // se agrega un LISTENER para que este pendiente de cuando se halla realizado la consulta y el servidor halla devuelto informacion
    xhr.addEventListener('load', e => {
      // con el self se hace referencia a la variable this
      let self = e.target;
    // estado
    // se guarda en un objeto por que hace mas facil el manejo cuando se "renderize" el HTML
      let respuesta = {
        status: self.status,
      // contenido de la respuesta del servidor
      // con este formato JSON.parse(self.response) la respuesta se mostrara en json y no como texto plano
        response: JSON.parse(self.response)
      };
      resolver(respuesta);
    });
    // errores propios de peticiones o conexion
    xhr.addEventListener('error', e => {
      let self = e.target;
      rechazar(self);
    });
    // se le pide al AJAX que envie el contenido
    xhr.send(obj);
  });
}


function $(elemento) {
    return document.getElementById(elemento);
}
