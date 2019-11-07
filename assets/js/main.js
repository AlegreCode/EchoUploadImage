function init(){
  const avatar = document.getElementById('file');
  avatar.addEventListener('change', function(e){
    const file = e.target.files[0];
    document.getElementById('image_preview').src = URL.createObjectURL(file);
  }, false);
}


document.addEventListener('DOMContentLoaded', init, false);