$.ajaxSetup({
    beforeSend: function(xhr, settings) {
        if (settings.type == 'POST' || settings.type == 'PUT' || settings.type == 'DELETE') {
            xhr.setRequestHeader('X-CSRF-Token', $('meta[name="csrf-token"]').attr('content'));
        }
    }
});

$(function() {
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
});
