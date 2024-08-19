let userData = document.getElementById("greeting")

let btn = document.querySelector("button")

btn.addEventListener("click", function (e) {
    let inputData = prompt("Put some text!")
    if (inputData === "a") {

        userData.textContent = "bullshit"
    }
    else{
        userData.textContent = inputData
    }
})

let arr = []
