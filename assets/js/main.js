function init(){
  const avatar = document.getElementById('file');
  avatar.addEventListener('change', function(e){
    const file = e.target.files[0];
    document.getElementById('image_preview').src = URL.createObjectURL(file);
  }, false);


 document.addEventListener("click", function (e) {
   elem = e.target
   if (elem.getAttribute("name") == "btn_delete") {
      e.preventDefault()
     Swal.fire({
      title: 'Are you sure?',
      text: "You won't be able to revert this!",
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
      if (result.value) {
        elem.parentNode.submit()
      }
    })
   }
  }, false);
}


document.addEventListener('DOMContentLoaded', init, false);