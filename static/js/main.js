$(document).ready(function() {
    $(".form").submit(function (e) {
        e.preventDefault()
        var fd = new FormData()
        var files = $(".file")[0].files;

        if (files.length >0 ){
            fd.append("file",files[0])

            $.ajax({
                url:"/",
                type:"post",
                data:fd,
                contentType:false,
                processData: false,
                success:function (response){

                },
                progress: function (e){
                    if (e.lengthComputable){
                        var pct =Math.round( (e.loaded/e.total)*100);
                        $(".progress").attr("value",pct)
                        $(".pct").text(`${pct}%`)
                        if (pct==100){
                            $(".pct").html("Done&check;")
                            $(".pct").css("color","mediumseagreen")
                        }

                    }
                }
            })
        }
        // alert("Form Submitted")
    })
})