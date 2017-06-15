let formLogin = $('form-login'),
    email = $('email'),
    password = $('password'),
    /*btnLogin = $('btnLogin'),*/
    mensajeLogin = $('mensaje-login');
// esto para que el formulario no se envie de forma nativa
formLogin.addEventListener('submit', e => {
    e.preventDefault();
  // se crea un objeto tipo js para enviarlo al api de loginlet obj = {
let obj = {
  email: email.value,
  password: password.value
};
// solicitud para usar la libreria que se creado
peticionAjax(formLogin.method, formLogin.action, JSON.stringify(obj))
.then(respuesta => {
  if (respuesta.status === 200) {
    mensajeLogin.textContent = 'Ingresaste';
    sessionStorage.setItem('token', respuesta.response.token);
    console.log(respuesta.response);
  } else {
    mensajeLogin.textContent = respuesta.response.message;
    console.log(respuesta.response);
  }
})
// confirmar errores en la comunicacion del servidor
.catch(error => {
  console.log(error);
})
});
