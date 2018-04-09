const PROXY_CONFIG = [
    {
        context: [
            "/api/**",
        ],
        target: "https://nhl-drafter.appspot.com",		
        secure: true,
        changeOrigin: true,
        cookieDomainRewrite: "localhost",
        https: true	
    }
]		
		
module.exports = PROXY_CONFIG;
