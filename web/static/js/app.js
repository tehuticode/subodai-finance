document.body.addEventListener('htmx:afterSwap', function(event) {
    if (event.detail.target.id === 'trade-result') {
        alert('Trade placed successfully!');
    }
});

// You can add more JavaScript functionality here as needed