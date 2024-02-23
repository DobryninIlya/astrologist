// Select all elements with the class 'colored_text'
let elements = document.querySelectorAll('.colored_text');

// Iterate over each element
elements.forEach((element) => {
    // Get the numerical content of the element
    let num = parseFloat(element.textContent);

    // Check if the content is a number
    if (!isNaN(num)) {
        // If the number is less than 0, set the color to red
        if (num < 0) {
            element.style.color = '#D32525';
        }
        // If the number is greater than 0, set the color to green
        else if (num > 0) {
            element.style.color = '#11AD33';
        }
    }
});