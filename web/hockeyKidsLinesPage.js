function submitForm(){
    let postData = {};
    $(":input").each(function() {
        postData[$(this).attr("name")] = parseInt($(this).val());
    });
    if (postData.numberOfPlayers % postData.numberOfPlayersPerLine == 0) {
        $("p").html("Sorry, you don't need help as you number of players can be divided by the number of players per line");
        return
    }
    if (postData.numberOfPlayers < 2 * postData.numberOfPlayersPerLine) {
        $("p").html("Sorry, you don't have enough players to build 2 lines");
        return
    }
    $.ajax({
        type: "post",
        url: window.location.pathname.startsWith('/web/hockeyKidsLinesPage.html') ? '../api' : '/api',
        data: JSON.stringify(postData),
        success: displayResult,
        error: function (data) {
            $("p").html("Sorry, an error occurred during the processing of your request, please try again later");
            $("form").remove()
        },
    });
}

function displayResult(data) {
    $("p").html("");
    let root = $(result);
    let innerDiv = "";
    let bestMatch = JSON.parse(data).bestMatch;
    if (bestMatch.length == 0) {
        $("p").html("Sorry, we didn't managed to fin a correct solution to your problem");
    } else {
        for (let i of bestMatch) {
            innerDiv += buildMatchView(i) + "\n"
        }
        root.html(innerDiv)
    }
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