import {Link, Stack, Typography} from "@mui/material";

export function Copyright() {
    return (
        <Stack direction="row" justifyContent="center" alignItems="center" sx={{flexGrow: 0}}>
            <Typography variant="body2" color="text.secondary" align="center">
                {'Copyright © '}
                <Link color="inherit" href="https://github.com/arikkfir">arikkfir</Link>
                {' ' + new Date().getFullYear()}.
            </Typography>
        </Stack>
    );
}
