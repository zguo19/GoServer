$(document).ready(function() {
    $('#submit').click(function (e) {
        var input = {};
        input.num1 = parseInt(document.getElementById('1').value);
        input.num2 = parseInt(document.getElementById('2').value);
        input.num3 = parseInt(document.getElementById('3').value);
        input.num4 = parseInt(document.getElementById('4').value);

        console.log(input);

        $.ajax({
            type: 'POST',
            data: JSON.stringify(input),
            url: '/api/math24',
            dataType: 'json',
            success: function (data) {
                console.log(data.Sum);
                document.getElementById('re').innerHTML = data.Sum;
            }
        })
    })
}) 