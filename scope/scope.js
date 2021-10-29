// difference between var keyword and let keyword 
//  var : global scope when declared outside function. local scope when declared inside function. 
//  let : block scope ( { } ). let keyword is introduced mainly because of
//  the function scope(var keyword) was confusing to cause most bugs in Javascript
// reference : https://www.tutorialspoint.com/difference-between-var-and-let-in-javascript


// example
var funcs = [];
// let's create 3 functions
for (var i = 0; i < 3; i++) {
  // and store them in funcs
  funcs[i] = function() {
    // each should log its value.
    console.log("My value: " + i);
  };
}
for (var j = 0; j < 3; j++) {
  // and now let's run each one to see
  funcs[j](); // value : 3, 3, 3 
}

let first = 'global';

function checkScope() {
    first = 'local';
    console.log(first);
}

checkScope(); // result : 'local'

// local variable vs global variable in Javascript
// Use local variable always when there is no need for global variable since
// other collaborators can use the variable name you thought it would work.
var vscope = 'global'; 
function fscope() { 
    // function firstly checks if there are local variables in function.
    // if exists, points to the local variable. 
    var vscope = 'local'; // local variable created
    vscope = 'local'; 
}
fscope(); 
console.log(`I am ${vscope}`);  // result : 'I am global' 

var a = 'hello';
var a = 'not hello'; // var keyword can be overwritten
console.log(a); 

let b = 'hello'; 
// let b = 'not hello'; // let keyword can not be overwritten 
console.log(b); 


// method to avoid of creating unnecessary global variables

// 1) anonymous function : creating a global variable without creating
// a global variable by creating a function and calling right away. 
(function myFn() {
    let myObject = { 
        "name":"hello function", 
        "age":2
    }
    console.log(`my name is: ${myObject.name}, my age is: ${myObject.age}`);
}()); 

// 2) creating an object and managing property 
const myObject = {
    "name_second":"good bye function", 
    "age_second":3
}

myObject.location_second = "USA";
console.log(myObject.location_second);


// Javascript only supports function scope coverage unlike Go and other
// language supports block scope coverage. Refer to scope.go for comparison
function eat() {
    let sushi = "sushi";
    console.log(sushi); 
}

eat(); 
// console.log(sushi); // should cause an undefined error 

// var number : global scope(outside function), let number : local scope
for (var number = 0; number < 3; number++) { 
    console.log(number); 
}
number = 55;
console.log(number); // result : 55