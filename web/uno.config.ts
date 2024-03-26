import { presetForms } from '@julr/unocss-preset-forms'
import {
    defineConfig,
    presetIcons,
    presetTypography,
    presetUno,
    presetWebFonts,
    transformerDirectives
} from 'unocss'

export default defineConfig({
    presets: [
        presetIcons({
            extraProperties: {
                display: 'inline-block',
                'vertical-align': 'middle'
            }
        }),
        presetForms(),
        presetUno(),
        presetTypography(),
        presetWebFonts({
            provider: 'google',
            fonts: {
                sans: 'Open Sans:100,200,300,400,500,600,700'
            }
        })
    ],
    transformers: [transformerDirectives()]
})
