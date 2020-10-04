// find input field
const rhymeElement = document.querySelector('#rhyme-input');

//add event listener
rhymeElement.oninput = handleInput;

// handle
function handleInput(event){
  console.log("Value", this.value)
}
