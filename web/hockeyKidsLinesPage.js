function submitForm(){
    let postUrl = "..?";
    let postData = {};
    $(":input").each(function() {
        postData[$(this).attr("name")] = $(this).attr("value");
    });
    console.log(postData)
    $.ajax({
        type: "post",
        url: postUrl,
        data: JSON.stringify(postData),
        success: function (data) {
            $("p").html("Thanks for your submit, your consent preferences have been saved");
            $("form").remove()
        },
        error: function (data) {
            $("p").html("An error occurred saving your consent preferences, please try again later");
            $("form").remove()
        },
    });
}