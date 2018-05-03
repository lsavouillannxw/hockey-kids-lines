function submitForm(){
    let postData = {};
    $(":input").each(function() {
        postData[$(this).attr("name")] = parseInt($(this).val());
    });
    console.log(postData)
    $.ajax({
        type: "post",
        url: "../api",
        data: JSON.stringify(postData),
        success: displayResult,
        error: function (data) {
            $("p").html("Sorry, an error occurred during the processing of your request, please try again later");
            $("form").remove()
        },
    });
}

function displayResult(data) {
    let root = $(result);
    let innerDiv = "";
    for (let i of JSON.parse(data).bestMatch) {
        innerDiv += buildMatchView(i) + "\n"
    }
    root.html(innerDiv)
}

function buildMatchView(data) {
    let result = ["<table border='1'>"];
    let i = 0;
    result.push(`<tr><th>Presence</th>`);
    for (let firstRow of data[0]) {
        result.push(`<th>${++i}</th>`);
    }
    i = 0;
    result.push("</tr>");
    for(let row of data) {
        result.push(`<tr><th>Player ${++i}</th>`);
        for(let cell of row.split("").reverse().join("")){
            result.push(`<td>${cell == '0' ? '': i}</td>`);
        }
        result.push("</tr>");
    }
    result.push("</table><br>");
    return result.join('\n');
}