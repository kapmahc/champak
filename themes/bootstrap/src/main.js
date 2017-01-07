import $ from 'jquery'
import marked from 'marked'

require("./main.css");

require("file-loader!./fail.png")
require("file-loader!./up.png")
require("file-loader!./down.png")
require("file-loader!./favicon.png")

$.ajaxSetup({
    beforeSend: function(xhr, settings) {
        if (settings.type == 'POST' || settings.type == 'PUT' || settings.type == 'DELETE') {
            xhr.setRequestHeader('X-CSRF-Token', $('meta[name="csrf-token"]').attr('content'));
        }
    }
});

$(function() {
  // ---------------------
  $('p.markdown').each(function(i, n){
    var txt = $(this).text();
    $(this).html(marked(txt));
  });
  // ---------------------
  $('a[data-method="delete"]').click(function(e) {
      e.preventDefault();
      if (confirm($(this).data('confirm'))) {
          $.ajax({
              url: $(this).attr('href'),
              type: 'DELETE',
              success: function(rst) {
                  window.location.href = $(this).data('next');
              }.bind(this)
          })
      }
  });
  // ---------------------
  $("img.votes").click(function(e){
    e.preventDefault();
    $.ajax({
        url: "/votes",
        type: 'POST',
        data: {
          point: $(this).data('point'),
          type: $(this).data('type'),
          id: $(this).data('id'),
        },
        success: function(rst) {
          alert(rst.message)
        }.bind(this)
    })
  });
  // ---------------------  
  // ---------------------
});
