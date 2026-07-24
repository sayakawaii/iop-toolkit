

def is_df_platform(model: str) -> bool:
    df_models = [
        "DFMB-B",
        "DFMB-C",
        "DFMB-D",
        "SFDB-A",
    ]
    return model in df_models

def is_brugal_platform(model: str) -> bool:
    brugal_models = [
        "BRUGAL-1U",
        "BRUGAL-2U",
        "BRUGAL-PLUS-1U",
        "BRUGAL-PLUS-2U",
    ]
    return model in brugal_models