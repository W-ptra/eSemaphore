function Header(){
    return (
        <header
            className="h-15 flex items-center bg-white z-30"
        >
            <div 
                className="w-15 bg-white z-30 h-full"
            />
            <div
                className="flex-1 flex items-center h-full bg-white"
                style={{ boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)' }}
            >
                <img 
                    src="/logo/eSemaphore_logo_white.webp" alt="" 
                    className="w-10 h-10"    
                />
                <h4 className="text-blue-400 font-bold ml-1 text-xl">
                    eSemaphore
                </h4>
            </div>
        </header>
    )
}

export default Header