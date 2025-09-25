import { Link } from 'react-router-dom'
import '../App.css'
import { easeInOut, motion } from "motion/react"
import { useState } from 'react'
import { delay } from 'motion'
import MotionLink from './MotionLink'

function Header() {
    // window location href is for new tab (re-renders otherwise) -> also for scripts. if not script, use <a target="_blank"></a>
    // useNavigate is for scripts
    // use Link or NavLink --> NavLink gives class="active"

    const [navBarOpen, setNavBarOpen] = useState<Boolean>(false)

    const similarVariant = {
        initial: {width: 100, opacity: 1},
        animate: {width: 0, opacity: 0},
        hover: {width: 100, opacity: 1}
    }

    const logoVariant = {
        initial: { x: 120 },
        animate: { x: 0, transition: {duration: 0.5, ease: easeInOut, delay: 0.4}},
        hover: { x: 120, transition: {duration: 0.3, ease: easeInOut}}
    }
    const ongsVariant = {
        initial: { x: 0, opacity: 1 },
        animate: { x: -100, opacity: 0, transition: {duration: 0.5, ease: easeInOut, delay: 0.4}},
        hover: { x: 0, opacity: 1, transition: {duration: 0.3, ease: easeInOut}}
    }

    return (
        <header 
            className='
                flex justify-between items-center relative 
                h-20 w-auto mx-[4%] mt-[10px] pb-[10px]
                border-b-1 border-white'
        >
            <MotionLink
                className='relative flex flex-row h-full text-4xl text-white font-electrolize' 
                to='/recsGenerator' 
                whileHover="hover" 
                initial="initial"
                animate="animate"
                key="logo-link"
                >
                 {/* text slide left */}
                <motion.div className='absolute top-[20px] right-20 h-full origin-right'
                    transition={{width: {duration: 0.5, ease: easeInOut, delay: 0.2}, opacity: {duration: 0.3, delay: 0.2}}}
                    variants={similarVariant}
                >
                    Similar
                </motion.div>
                
                {/* img + txt right */}
                <motion.div 
                    // TODO bg same colour as header
                    className='flex justify-center items-center h-full z-0 bg-[#141b2a]'
                    variants={logoVariant}
                >
                    <img src='/similarSongsLogo.png' className='h-full'/>
                    <motion.div className='w-30 justify-left'
                        variants={ongsVariant}
                    >ongs</motion.div>
                </motion.div>
               
            </MotionLink>
            
            <button className='flex justify-center items-center flex-col w-20 h-20 group relative'
                onClick={() => setNavBarOpen(prev => !prev)}
            >
                {/* hamburg */}
                <div className={`bg-white w-[50%] h-1 m-1 transition-transform duration-600 ${navBarOpen ? "-rotate-270 translate-y-3 translate-x-2" : ""}`}/>
                <div className={`bg-white w-[50%] h-1 m-1 transition-transform duration-600 ${navBarOpen ? "-rotate-270 -translate-x-2" : ""}`}/>
                <div className={`bg-white w-[50%] h-1 m-1 transition-transform duration-600 ${navBarOpen ? "-rotate-270 -translate-y-3" : ""}`}/>

                {/* navbar */}
                <nav></nav>
            </button>
        </header>

    )
}

export default Header