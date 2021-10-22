const bookshelf = {
    'category':{"horror":10, "romance":7, "comedy":5}, 
    'count':function(){
        // 'this' keyword in Javascript is a resevered keyword, 
        //  which points to the object that it belong to. 
        let total = 0;
        for (key in this.category){ 
            total += this.category[key];
        }
        console.log(total); 
    }
}

// language comparision 1 : struct in Go
bookshelf.category; 

// language comparision 2 : method in Go
bookshelf.count();